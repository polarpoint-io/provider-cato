/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// nestedRuleID is the external name configuration for the policy rule
// resources. The Cato Management Application assigns the rule an ID on create,
// but terraform-provider-cato is built on terraform-plugin-framework and keeps
// that ID inside the "rule" object rather than exposing a top-level "id"
// attribute, so IdentifierFromProvider would read an empty external name.
var nestedRuleID = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn: func(tfstate map[string]any) (string, error) {
		rule, ok := tfstate["rule"].(map[string]any)
		if !ok {
			return "", nil
		}
		id, ok := rule["id"].(string)
		if !ok {
			return "", nil
		}
		return id, nil
	},
	GetIDFn: func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		return externalName, nil
	},
	DisableNameInitializer: true,
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// The Cato Management Application assigns every object an opaque ID on create,
// so nothing here is named by the user: the external name annotation carries
// the CMA ID and metadata.name is free to be whatever you like.
var ExternalNameConfigs = map[string]config.ExternalName{
	// access
	"cato_app_connector":            config.IdentifierFromProvider,
	"cato_private_access_policy":    config.IdentifierFromProvider,
	"cato_private_access_rule":      config.IdentifierFromProvider,
	"cato_private_access_rule_bulk": config.IdentifierFromProvider,
	"cato_private_app":              config.IdentifierFromProvider,

	// admin
	"cato_account":       config.IdentifierFromProvider,
	"cato_admin":         config.IdentifierFromProvider,
	"cato_group":         config.IdentifierFromProvider,
	"cato_group_members": config.IdentifierFromProvider,

	// policy
	"cato_app_tenant_restriction_rule":    config.IdentifierFromProvider,
	"cato_app_tenant_restriction_section": config.IdentifierFromProvider,
	"cato_application_control_policy":     config.IdentifierFromProvider,
	"cato_application_control_rule":       config.IdentifierFromProvider,
	"cato_application_control_section":    config.IdentifierFromProvider,
	"cato_if_rule":                        nestedRuleID,
	"cato_if_section":                     config.IdentifierFromProvider,
	"cato_if_sub_policy":                  config.IdentifierFromProvider,
	"cato_socket_lan_firewall_rule":       nestedRuleID,
	"cato_socket_lan_network_rule":        nestedRuleID,
	"cato_socket_lan_section":             config.IdentifierFromProvider,
	"cato_tls_rule":                       config.IdentifierFromProvider,
	"cato_tls_section":                    config.IdentifierFromProvider,
	"cato_wf_rule":                        nestedRuleID,
	"cato_wf_section":                     config.IdentifierFromProvider,
	"cato_wf_sub_policy":                  config.IdentifierFromProvider,
	"cato_wnw_rule":                       nestedRuleID,
	"cato_wnw_section":                    config.IdentifierFromProvider,

	// site
	"cato_bgp_peer":                 config.IdentifierFromProvider,
	"cato_ipsec_site":               config.IdentifierFromProvider,
	"cato_lan_interface":            config.IdentifierFromProvider,
	"cato_lan_interface_lag_member": config.IdentifierFromProvider,
	"cato_license":                  config.IdentifierFromProvider,
	"cato_network_range":            config.IdentifierFromProvider,
	"cato_socket_site":              config.IdentifierFromProvider,
	"cato_static_host":              config.IdentifierFromProvider,
	"cato_wan_interface":            config.IdentifierFromProvider,
}

// SkippedResources lists the Terraform resources deliberately left out of the
// generated API.
//
// These are bulk helpers rather than objects: the cato_bulk_*_move_rule
// resources reorder rules that other resources own, and cato_global_ip_ranges
// replaces an account-wide list wholesale. None of them has an identity of its
// own in Terraform state, so there is nothing for a managed resource's external
// name to point at and each reconcile would re-run the operation.
var SkippedResources = []string{
	"cato_bulk_if_move_rule$",
	"cato_bulk_tls_move_rule$",
	"cato_bulk_wf_move_rule$",
	"cato_bulk_wnw_move_rule$",
	"cato_global_ip_ranges$",
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
