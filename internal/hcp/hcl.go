package hcp

import (
	"context"
	"fmt"

	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"
	imgds "github.com/hashicorp/packer/datasource/hcp-packer-image"
	iterds "github.com/hashicorp/packer/datasource/hcp-packer-iteration"
	"github.com/hashicorp/packer/hcl2template"
	"github.com/hashicorp/packer/internal/registry"
	"github.com/hashicorp/packer/internal/registry/env"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

const (
	// Known HCP Packer Image Datasource, whose id is the SourceImageId for some build.
	hcpImageDatasourceType     string = "hcp-packer-image"
	hcpIterationDatasourceType string = "hcp-packer-iteration"
	buildLabel                 string = "build"
)

// hclOrchestrator is a HCP handler made for handling HCL configurations
type hclOrchestrator struct {
	configuration *hcl2template.PackerConfig
	bucket        *registry.Bucket
}

// PopulateIteration creates the metadata on HCP for a build
func (h *hclOrchestrator) PopulateIteration(ctx context.Context) error {
	for _, build := range h.configuration.Builds {
		build.HCPPackerRegistry.WriteToBucketConfig(h.bucket)

		// If at this point the bucket.Slug is still empty,
		// last try is to use the build.Name if present
		if h.bucket.Slug == "" && build.Name != "" {
			h.bucket.Slug = build.Name
		}

		// If the description is empty, use the one from the build block
		if h.bucket.Description == "" {
			h.bucket.Description = build.Description
		}

		for _, source := range build.Sources {
			h.bucket.RegisterBuildForComponent(source.String())
		}
	}

	err := h.getInfoFromDataSources()
	if err != nil {
		return err
	}

	err = h.bucket.Initialize(ctx)
	if err != nil {
		return err
	}

	err = h.bucket.PopulateIteration(ctx)
	if err != nil {
		return err
	}

	iterationID := h.bucket.Iteration.ID

	h.configuration.HCPVars["iterationID"] = cty.StringVal(iterationID)

	return nil
}

func (h *hclOrchestrator) getInfoFromDataSources() error {
	vals, diags := h.configuration.Datasources.Values()
	if diags != nil {
		return diags
	}

	imageDS, imageOK := vals[hcpImageDatasourceType]
	iterDS, iterOK := vals[hcpIterationDatasourceType]

	// If we don't have any image or iteration defined, we can return directly
	if !imageOK && !iterOK {
		return nil
	}

	iterations := map[string]iterds.DatasourceOutput{}

	if iterOK {
		hcpData := map[string]cty.Value{}
		err := gocty.FromCtyValue(iterDS, &hcpData)
		if err != nil {
			return fmt.Errorf("Failed to decode hcp_packer_iteration datasources: %s", err)
		}

		for k, v := range hcpData {
			iterVals := v.AsValueMap()
			dso := iterValueToDSOutput(iterVals)
			iterations[k] = dso
		}
	}

	images := map[string]imgds.DatasourceOutput{}

	if imageOK {
		hcpData := map[string]cty.Value{}
		err := gocty.FromCtyValue(imageDS, &hcpData)
		if err != nil {
			return fmt.Errorf("Failed to decode hcp_packer_image datasources: %s", err)
		}

		for k, v := range hcpData {
			imageVals := v.AsValueMap()
			dso := imageValueToDSOutput(imageVals)
			images[k] = dso
		}
	}

	for _, img := range images {
		sourceIteration := registry.ParentIteration{}

		sourceIteration.IterationID = img.IterationID

		if img.ChannelID != "" {
			sourceIteration.ChannelID = img.ChannelID
		} else {
			for _, it := range iterations {
				if it.ID == img.IterationID {
					sourceIteration.ChannelID = it.ChannelID
					break
				}
			}
		}

		h.bucket.SourceImagesToParentIterations[img.ID] = sourceIteration
	}

	return nil
}

// BuildStart is invoked when one build for the configuration is starting to be processed
func (h *hclOrchestrator) BuildStart(ctx context.Context, buildName string) error {
	return h.bucket.BuildStart(ctx, buildName)
}

// BuildDone is invoked when one build for the configuration has finished
func (h *hclOrchestrator) BuildDone(
	ctx context.Context,
	buildName string,
	artifacts []sdkpacker.Artifact,
	buildErr error,
) ([]sdkpacker.Artifact, error) {
	return h.bucket.BuildDone(ctx, buildName, artifacts, buildErr)
}

func newHCLOrchestrator(config *hcl2template.PackerConfig) (Orchestrator, error) {
	hasHCP := false

	for _, build := range config.Builds {
		if build.HCPPackerRegistry != nil {
			hasHCP = true
		}
	}

	if env.IsPARDisabled() {
		hasHCP = false
	}

	if !hasHCP {
		return newNoopHandler(), nil
	}

	if len(config.Builds) > 1 {
		return nil, fmt.Errorf("For Packer Registry enabled builds, only one " + buildLabel +
			" block can be defined. Please remove any additional " + buildLabel +
			" block(s). If this " + buildLabel + " is not meant for the Packer registry please " +
			"clear any HCP_PACKER_* environment variables or `hcp_packer_registry` block.")
	}

	bucket, err := registry.NewBucketWithIteration(registry.IterationOptions{
		TemplateBaseDir: config.Basedir,
	})

	if err != nil {
		return nil, fmt.Errorf(
			"Unable to create a valid bucket object for HCP Packer Registry: %s",
			err)
	}

	bucket.LoadDefaultSettingsFromEnv()

	return &hclOrchestrator{
		configuration: config,
		bucket:        bucket,
	}, nil
}

func imageValueToDSOutput(imageVal map[string]cty.Value) imgds.DatasourceOutput {
	dso := imgds.DatasourceOutput{}
	for k, v := range imageVal {
		switch k {
		case "id":
			dso.ID = v.AsString()
		case "region":
			dso.Region = v.AsString()
		case "labels":
			labels := map[string]string{}
			lbls := v.AsValueMap()
			for k, v := range lbls {
				labels[k] = v.AsString()
			}
			dso.Labels = labels
		case "packer_run_uuid":
			dso.PackerRunUUID = v.AsString()
		case "channel_id":
			dso.ChannelID = v.AsString()
		case "iteration_id":
			dso.IterationID = v.AsString()
		case "build_id":
			dso.BuildID = v.AsString()
		case "created_at":
			dso.CreatedAt = v.AsString()
		case "component_type":
			dso.ComponentType = v.AsString()
		case "cloud_provider":
			dso.CloudProvider = v.AsString()
		}
	}

	return dso
}

func iterValueToDSOutput(iterVal map[string]cty.Value) iterds.DatasourceOutput {
	dso := iterds.DatasourceOutput{}
	for k, v := range iterVal {
		switch k {
		case "author_id":
			dso.AuthorID = v.AsString()
		case "bucket_name":
			dso.BucketName = v.AsString()
		case "complete":
			// For all intents and purposes, cty.Value.True() acts
			// like a AsBool() would.
			dso.Complete = v.True()
		case "created_at":
			dso.CreatedAt = v.AsString()
		case "fingerprint":
			dso.Fingerprint = v.AsString()
		case "id":
			dso.ID = v.AsString()
		case "incremental_version":
			// Maybe when cty provides a good way to AsInt() a cty.Value
			// we can consider implementing this.
		case "updated_at":
			dso.UpdatedAt = v.AsString()
		case "channel_id":
			dso.ChannelID = v.AsString()
		}
	}
	return dso
}
