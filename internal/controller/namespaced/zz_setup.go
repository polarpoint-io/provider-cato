// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appconnector "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/access/appconnector"
	privateaccesspolicy "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/access/privateaccesspolicy"
	privateaccessrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/access/privateaccessrule"
	privateaccessrulebulk "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/access/privateaccessrulebulk"
	privateapp "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/access/privateapp"
	account "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/admin/account"
	admin "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/admin/admin"
	group "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/admin/group"
	groupmembers "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/admin/groupmembers"
	applicationcontrolpolicy "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/applicationcontrolpolicy"
	applicationcontrolrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/applicationcontrolrule"
	applicationcontrolsection "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/applicationcontrolsection"
	apptenantrestrictionrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/apptenantrestrictionrule"
	apptenantrestrictionsection "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/apptenantrestrictionsection"
	ifrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/ifrule"
	ifsection "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/ifsection"
	ifsubpolicy "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/ifsubpolicy"
	socketlanfirewallrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/socketlanfirewallrule"
	socketlannetworkrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/socketlannetworkrule"
	socketlansection "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/socketlansection"
	tlsrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/tlsrule"
	tlssection "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/tlssection"
	wfrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/wfrule"
	wfsection "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/wfsection"
	wfsubpolicy "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/wfsubpolicy"
	wnwrule "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/wnwrule"
	wnwsection "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/policy/wnwsection"
	providerconfig "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/providerconfig"
	bgppeer "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/bgppeer"
	ipsecsite "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/ipsecsite"
	laninterface "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/laninterface"
	laninterfacelagmember "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/laninterfacelagmember"
	license "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/license"
	networkrange "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/networkrange"
	socketsite "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/socketsite"
	statichost "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/statichost"
	waninterface "github.com/polarpoint-io/provider-cato/internal/controller/namespaced/site/waninterface"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnector.Setup,
		privateaccesspolicy.Setup,
		privateaccessrule.Setup,
		privateaccessrulebulk.Setup,
		privateapp.Setup,
		account.Setup,
		admin.Setup,
		group.Setup,
		groupmembers.Setup,
		applicationcontrolpolicy.Setup,
		applicationcontrolrule.Setup,
		applicationcontrolsection.Setup,
		apptenantrestrictionrule.Setup,
		apptenantrestrictionsection.Setup,
		ifrule.Setup,
		ifsection.Setup,
		ifsubpolicy.Setup,
		socketlanfirewallrule.Setup,
		socketlannetworkrule.Setup,
		socketlansection.Setup,
		tlsrule.Setup,
		tlssection.Setup,
		wfrule.Setup,
		wfsection.Setup,
		wfsubpolicy.Setup,
		wnwrule.Setup,
		wnwsection.Setup,
		providerconfig.Setup,
		bgppeer.Setup,
		ipsecsite.Setup,
		laninterface.Setup,
		laninterfacelagmember.Setup,
		license.Setup,
		networkrange.Setup,
		socketsite.Setup,
		statichost.Setup,
		waninterface.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnector.SetupGated,
		privateaccesspolicy.SetupGated,
		privateaccessrule.SetupGated,
		privateaccessrulebulk.SetupGated,
		privateapp.SetupGated,
		account.SetupGated,
		admin.SetupGated,
		group.SetupGated,
		groupmembers.SetupGated,
		applicationcontrolpolicy.SetupGated,
		applicationcontrolrule.SetupGated,
		applicationcontrolsection.SetupGated,
		apptenantrestrictionrule.SetupGated,
		apptenantrestrictionsection.SetupGated,
		ifrule.SetupGated,
		ifsection.SetupGated,
		ifsubpolicy.SetupGated,
		socketlanfirewallrule.SetupGated,
		socketlannetworkrule.SetupGated,
		socketlansection.SetupGated,
		tlsrule.SetupGated,
		tlssection.SetupGated,
		wfrule.SetupGated,
		wfsection.SetupGated,
		wfsubpolicy.SetupGated,
		wnwrule.SetupGated,
		wnwsection.SetupGated,
		providerconfig.SetupGated,
		bgppeer.SetupGated,
		ipsecsite.SetupGated,
		laninterface.SetupGated,
		laninterfacelagmember.SetupGated,
		license.SetupGated,
		networkrange.SetupGated,
		socketsite.SetupGated,
		statichost.SetupGated,
		waninterface.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		appconnector.SetupWebhookWithManager,
		privateaccesspolicy.SetupWebhookWithManager,
		privateaccessrule.SetupWebhookWithManager,
		privateaccessrulebulk.SetupWebhookWithManager,
		privateapp.SetupWebhookWithManager,
		account.SetupWebhookWithManager,
		admin.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		groupmembers.SetupWebhookWithManager,
		applicationcontrolpolicy.SetupWebhookWithManager,
		applicationcontrolrule.SetupWebhookWithManager,
		applicationcontrolsection.SetupWebhookWithManager,
		apptenantrestrictionrule.SetupWebhookWithManager,
		apptenantrestrictionsection.SetupWebhookWithManager,
		ifrule.SetupWebhookWithManager,
		ifsection.SetupWebhookWithManager,
		ifsubpolicy.SetupWebhookWithManager,
		socketlanfirewallrule.SetupWebhookWithManager,
		socketlannetworkrule.SetupWebhookWithManager,
		socketlansection.SetupWebhookWithManager,
		tlsrule.SetupWebhookWithManager,
		tlssection.SetupWebhookWithManager,
		wfrule.SetupWebhookWithManager,
		wfsection.SetupWebhookWithManager,
		wfsubpolicy.SetupWebhookWithManager,
		wnwrule.SetupWebhookWithManager,
		wnwsection.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		bgppeer.SetupWebhookWithManager,
		ipsecsite.SetupWebhookWithManager,
		laninterface.SetupWebhookWithManager,
		laninterfacelagmember.SetupWebhookWithManager,
		license.SetupWebhookWithManager,
		networkrange.SetupWebhookWithManager,
		socketsite.SetupWebhookWithManager,
		statichost.SetupWebhookWithManager,
		waninterface.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
