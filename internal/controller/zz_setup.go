/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	databasev1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/dbaas/databasev1"
	datastorev1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/dbaas/datastorev1"
	extensionv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/dbaas/extensionv1"
	grantv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/dbaas/grantv1"
	prometheusmetrictokenv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/dbaas/prometheusmetrictokenv1"
	userv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/dbaas/userv1"
	domainv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/domains/domainv1"
	recordv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/domains/recordv1"
	clusterv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/mks/clusterv1"
	nodegroupv1 "github.com/ktsstudio/provider-jet-selectel/internal/controller/mks/nodegroupv1"
	providerconfig "github.com/ktsstudio/provider-jet-selectel/internal/controller/providerconfig"
	crossregionsubnetv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/crossregionsubnetv2"
	floatingipv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/floatingipv2"
	keypairv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/keypairv2"
	licensev2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/licensev2"
	projectv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/projectv2"
	rolev2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/rolev2"
	subnetv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/subnetv2"
	tokenv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/tokenv2"
	userv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/userv2"
	vrrpsubnetv2 "github.com/ktsstudio/provider-jet-selectel/internal/controller/vpc/vrrpsubnetv2"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		databasev1.Setup,
		datastorev1.Setup,
		extensionv1.Setup,
		grantv1.Setup,
		prometheusmetrictokenv1.Setup,
		userv1.Setup,
		domainv1.Setup,
		recordv1.Setup,
		clusterv1.Setup,
		nodegroupv1.Setup,
		providerconfig.Setup,
		crossregionsubnetv2.Setup,
		floatingipv2.Setup,
		keypairv2.Setup,
		licensev2.Setup,
		projectv2.Setup,
		rolev2.Setup,
		subnetv2.Setup,
		tokenv2.Setup,
		userv2.Setup,
		vrrpsubnetv2.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
