// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"context"

	calicov1alpha1 "github.com/gardener/gardener-extension-networking-calico/pkg/apis/calico/v1alpha1"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (a *actuator) updateProviderStatus(
	ctx context.Context,
	network *extensionsv1alpha1.Network,
	config *calicov1alpha1.NetworkConfig,
) error {
	status, err := a.ComputeNetworkStatus(config)
	if err != nil {
		return err
	}

	patch := client.MergeFrom(network.DeepCopy())
	network.Status.ProviderStatus = &runtime.RawExtension{Object: status}
	network.Status.LastOperation = extensionscontroller.LastOperation(gardencorev1beta1.LastOperationTypeReconcile,
		gardencorev1beta1.LastOperationStateSucceeded,
		100,
		"Calico was configured successfully",
	)
	return a.client.Status().Patch(ctx, network, patch)
}

func (a *actuator) ComputeNetworkStatus(networkConfig *calicov1alpha1.NetworkConfig) (*calicov1alpha1.NetworkStatus, error) {
	var (
		status = &calicov1alpha1.NetworkStatus{
			TypeMeta: StatusTypeMeta,
		}
	)

	return status, nil
}
