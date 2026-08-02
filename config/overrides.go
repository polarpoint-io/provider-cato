/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// groupKindOverrides maps a Terraform resource name onto the short API group
// and kind its CRD should use.
//
// Upjet derives the group from the second word of the Terraform resource name,
// which here would produce a group per policy type and per interface type (if,
// wf, wnw, tls, socket, app, application, private, lan, wan, bgp, network,
// static). The resources are instead grouped by what they configure: site,
// policy, access and admin.
var groupKindOverrides = map[string]struct {
	group string
	kind  string
}{
	// access.cato.crossplane.io
	"cato_app_connector":            {group: "access", kind: "AppConnector"},
	"cato_private_access_policy":    {group: "access", kind: "PrivateAccessPolicy"},
	"cato_private_access_rule":      {group: "access", kind: "PrivateAccessRule"},
	"cato_private_access_rule_bulk": {group: "access", kind: "PrivateAccessRuleBulk"},
	"cato_private_app":              {group: "access", kind: "PrivateApp"},

	// admin.cato.crossplane.io
	"cato_account":       {group: "admin", kind: "Account"},
	"cato_admin":         {group: "admin", kind: "Admin"},
	"cato_group":         {group: "admin", kind: "Group"},
	"cato_group_members": {group: "admin", kind: "GroupMembers"},

	// policy.cato.crossplane.io
	"cato_app_tenant_restriction_rule":    {group: "policy", kind: "AppTenantRestrictionRule"},
	"cato_app_tenant_restriction_section": {group: "policy", kind: "AppTenantRestrictionSection"},
	"cato_application_control_policy":     {group: "policy", kind: "ApplicationControlPolicy"},
	"cato_application_control_rule":       {group: "policy", kind: "ApplicationControlRule"},
	"cato_application_control_section":    {group: "policy", kind: "ApplicationControlSection"},
	"cato_if_rule":                        {group: "policy", kind: "IfRule"},
	"cato_if_section":                     {group: "policy", kind: "IfSection"},
	"cato_if_sub_policy":                  {group: "policy", kind: "IfSubPolicy"},
	"cato_socket_lan_firewall_rule":       {group: "policy", kind: "SocketLanFirewallRule"},
	"cato_socket_lan_network_rule":        {group: "policy", kind: "SocketLanNetworkRule"},
	"cato_socket_lan_section":             {group: "policy", kind: "SocketLanSection"},
	"cato_tls_rule":                       {group: "policy", kind: "TlsRule"},
	"cato_tls_section":                    {group: "policy", kind: "TlsSection"},
	"cato_wf_rule":                        {group: "policy", kind: "WfRule"},
	"cato_wf_section":                     {group: "policy", kind: "WfSection"},
	"cato_wf_sub_policy":                  {group: "policy", kind: "WfSubPolicy"},
	"cato_wnw_rule":                       {group: "policy", kind: "WnwRule"},
	"cato_wnw_section":                    {group: "policy", kind: "WnwSection"},

	// site.cato.crossplane.io
	"cato_bgp_peer":                 {group: "site", kind: "BgpPeer"},
	"cato_ipsec_site":               {group: "site", kind: "IpsecSite"},
	"cato_lan_interface":            {group: "site", kind: "LanInterface"},
	"cato_lan_interface_lag_member": {group: "site", kind: "LanInterfaceLagMember"},
	"cato_license":                  {group: "site", kind: "License"},
	"cato_network_range":            {group: "site", kind: "NetworkRange"},
	"cato_socket_site":              {group: "site", kind: "SocketSite"},
	"cato_static_host":              {group: "site", kind: "StaticHost"},
	"cato_wan_interface":            {group: "site", kind: "WanInterface"},
}

// GroupKindOverrides returns a ResourceOption that applies the group and kind
// overrides in groupKindOverrides.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if o, ok := groupKindOverrides[r.Name]; ok {
			r.ShortGroup = o.group
			r.Kind = o.kind
		}
	}
}
