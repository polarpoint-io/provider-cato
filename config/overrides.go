/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// The short API groups this provider publishes, under cato.crossplane.io.
const (
	groupAccess = "access"
	groupAdmin  = "admin"
	groupPolicy = "policy"
	groupSite   = "site"
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
	"cato_app_connector":            {group: groupAccess, kind: "AppConnector"},
	"cato_private_access_policy":    {group: groupAccess, kind: "PrivateAccessPolicy"},
	"cato_private_access_rule":      {group: groupAccess, kind: "PrivateAccessRule"},
	"cato_private_access_rule_bulk": {group: groupAccess, kind: "PrivateAccessRuleBulk"},
	"cato_private_app":              {group: groupAccess, kind: "PrivateApp"},

	// admin.cato.crossplane.io
	"cato_account":       {group: groupAdmin, kind: "Account"},
	"cato_admin":         {group: groupAdmin, kind: "Admin"},
	"cato_group":         {group: groupAdmin, kind: "Group"},
	"cato_group_members": {group: groupAdmin, kind: "GroupMembers"},

	// policy.cato.crossplane.io
	"cato_app_tenant_restriction_rule":    {group: groupPolicy, kind: "AppTenantRestrictionRule"},
	"cato_app_tenant_restriction_section": {group: groupPolicy, kind: "AppTenantRestrictionSection"},
	"cato_application_control_policy":     {group: groupPolicy, kind: "ApplicationControlPolicy"},
	"cato_application_control_rule":       {group: groupPolicy, kind: "ApplicationControlRule"},
	"cato_application_control_section":    {group: groupPolicy, kind: "ApplicationControlSection"},
	"cato_if_rule":                        {group: groupPolicy, kind: "IfRule"},
	"cato_if_section":                     {group: groupPolicy, kind: "IfSection"},
	"cato_if_sub_policy":                  {group: groupPolicy, kind: "IfSubPolicy"},
	"cato_socket_lan_firewall_rule":       {group: groupPolicy, kind: "SocketLanFirewallRule"},
	"cato_socket_lan_network_rule":        {group: groupPolicy, kind: "SocketLanNetworkRule"},
	"cato_socket_lan_section":             {group: groupPolicy, kind: "SocketLanSection"},
	"cato_tls_rule":                       {group: groupPolicy, kind: "TlsRule"},
	"cato_tls_section":                    {group: groupPolicy, kind: "TlsSection"},
	"cato_wf_rule":                        {group: groupPolicy, kind: "WfRule"},
	"cato_wf_section":                     {group: groupPolicy, kind: "WfSection"},
	"cato_wf_sub_policy":                  {group: groupPolicy, kind: "WfSubPolicy"},
	"cato_wnw_rule":                       {group: groupPolicy, kind: "WnwRule"},
	"cato_wnw_section":                    {group: groupPolicy, kind: "WnwSection"},

	// site.cato.crossplane.io
	"cato_bgp_peer":                 {group: groupSite, kind: "BgpPeer"},
	"cato_ipsec_site":               {group: groupSite, kind: "IpsecSite"},
	"cato_lan_interface":            {group: groupSite, kind: "LanInterface"},
	"cato_lan_interface_lag_member": {group: groupSite, kind: "LanInterfaceLagMember"},
	"cato_license":                  {group: groupSite, kind: "License"},
	"cato_network_range":            {group: groupSite, kind: "NetworkRange"},
	"cato_socket_site":              {group: groupSite, kind: "SocketSite"},
	"cato_static_host":              {group: groupSite, kind: "StaticHost"},
	"cato_wan_interface":            {group: groupSite, kind: "WanInterface"},
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
