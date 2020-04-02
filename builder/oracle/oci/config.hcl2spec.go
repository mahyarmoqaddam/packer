// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package oci

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string                           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string                           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool                             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool                             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string                           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string                 `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string                          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Type                      *string                           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string                           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string                           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int                              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string                           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string                           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string                           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string                           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool                             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string                           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool                             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string                           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool                             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool                             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int                              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string                           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int                              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool                             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string                           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string                           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionInteractive     *bool                             `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string                           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string                           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string                           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int                              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string                           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string                           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string                           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string                           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string                          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string                          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte                            `mapstructure:"ssh_public_key" cty:"ssh_public_key"`
	SSHPrivateKey             []byte                            `mapstructure:"ssh_private_key" cty:"ssh_private_key"`
	WinRMUser                 *string                           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string                           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string                           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int                              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string                           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool                             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool                             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool                             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	InstancePrincipals        *bool                             `mapstructure:"use_instance_principals" cty:"use_instance_principals"`
	AccessCfgFile             *string                           `mapstructure:"access_cfg_file" cty:"access_cfg_file"`
	AccessCfgFileAccount      *string                           `mapstructure:"access_cfg_file_account" cty:"access_cfg_file_account"`
	UserID                    *string                           `mapstructure:"user_ocid" cty:"user_ocid"`
	TenancyID                 *string                           `mapstructure:"tenancy_ocid" cty:"tenancy_ocid"`
	Region                    *string                           `mapstructure:"region" cty:"region"`
	Fingerprint               *string                           `mapstructure:"fingerprint" cty:"fingerprint"`
	KeyFile                   *string                           `mapstructure:"key_file" cty:"key_file"`
	PassPhrase                *string                           `mapstructure:"pass_phrase" cty:"pass_phrase"`
	UsePrivateIP              *bool                             `mapstructure:"use_private_ip" cty:"use_private_ip"`
	AvailabilityDomain        *string                           `mapstructure:"availability_domain" cty:"availability_domain"`
	CompartmentID             *string                           `mapstructure:"compartment_ocid" cty:"compartment_ocid"`
	BaseImageID               *string                           `mapstructure:"base_image_ocid" cty:"base_image_ocid"`
	Shape                     *string                           `mapstructure:"shape" cty:"shape"`
	ImageName                 *string                           `mapstructure:"image_name" cty:"image_name"`
	InstanceName              *string                           `mapstructure:"instance_name" cty:"instance_name"`
	Metadata                  map[string]string                 `mapstructure:"metadata" cty:"metadata"`
	UserData                  *string                           `mapstructure:"user_data" cty:"user_data"`
	UserDataFile              *string                           `mapstructure:"user_data_file" cty:"user_data_file"`
	SubnetID                  *string                           `mapstructure:"subnet_ocid" cty:"subnet_ocid"`
	Tags                      map[string]string                 `mapstructure:"tags" cty:"tags"`
	DefinedTags               map[string]map[string]interface{} `mapstructure:"defined_tags" cty:"defined_tags"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"use_instance_principals":      &hcldec.AttrSpec{Name: "use_instance_principals", Type: cty.Bool, Required: false},
		"access_cfg_file":              &hcldec.AttrSpec{Name: "access_cfg_file", Type: cty.String, Required: false},
		"access_cfg_file_account":      &hcldec.AttrSpec{Name: "access_cfg_file_account", Type: cty.String, Required: false},
		"user_ocid":                    &hcldec.AttrSpec{Name: "user_ocid", Type: cty.String, Required: false},
		"tenancy_ocid":                 &hcldec.AttrSpec{Name: "tenancy_ocid", Type: cty.String, Required: false},
		"region":                       &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"fingerprint":                  &hcldec.AttrSpec{Name: "fingerprint", Type: cty.String, Required: false},
		"key_file":                     &hcldec.AttrSpec{Name: "key_file", Type: cty.String, Required: false},
		"pass_phrase":                  &hcldec.AttrSpec{Name: "pass_phrase", Type: cty.String, Required: false},
		"use_private_ip":               &hcldec.AttrSpec{Name: "use_private_ip", Type: cty.Bool, Required: false},
		"availability_domain":          &hcldec.AttrSpec{Name: "availability_domain", Type: cty.String, Required: false},
		"compartment_ocid":             &hcldec.AttrSpec{Name: "compartment_ocid", Type: cty.String, Required: false},
		"base_image_ocid":              &hcldec.AttrSpec{Name: "base_image_ocid", Type: cty.String, Required: false},
		"shape":                        &hcldec.AttrSpec{Name: "shape", Type: cty.String, Required: false},
		"image_name":                   &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"instance_name":                &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"metadata":                     &hcldec.BlockAttrsSpec{TypeName: "metadata", ElementType: cty.String, Required: false},
		"user_data":                    &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":               &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"subnet_ocid":                  &hcldec.AttrSpec{Name: "subnet_ocid", Type: cty.String, Required: false},
		"tags":                         &hcldec.BlockAttrsSpec{TypeName: "tags", ElementType: cty.String, Required: false},
		"defined_tags":                 &hcldec.BlockAttrsSpec{TypeName: "defined_tags", ElementType: cty.String, Required: false},
	}
	return s
}
