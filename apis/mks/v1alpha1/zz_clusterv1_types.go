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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KubeAPIIP *string `json:"kubeApiIp,omitempty" tf:"kube_api_ip,omitempty"`

	MaintenanceWindowEnd *string `json:"maintenanceWindowEnd,omitempty" tf:"maintenance_window_end,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ClusterV1Parameters struct {

	// +kubebuilder:validation:Optional
	AdmissionControllers []*string `json:"admissionControllers,omitempty" tf:"admission_controllers,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAutorepair *bool `json:"enableAutorepair,omitempty" tf:"enable_autorepair,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePatchVersionAutoUpgrade *bool `json:"enablePatchVersionAutoUpgrade,omitempty" tf:"enable_patch_version_auto_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePodSecurityPolicy *bool `json:"enablePodSecurityPolicy,omitempty" tf:"enable_pod_security_policy,omitempty"`

	// +kubebuilder:validation:Optional
	FeatureGates []*string `json:"featureGates,omitempty" tf:"feature_gates,omitempty"`

	// +kubebuilder:validation:Required
	KubeVersion *string `json:"kubeVersion" tf:"kube_version,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindowStart *string `json:"maintenanceWindowStart,omitempty" tf:"maintenance_window_start,omitempty"`

	// +crossplane:generate:reference:type=github.com/ktsstudio/provider-jet-selectel/apis/mks/v1alpha1.ClusterV1
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Zonal *bool `json:"zonal,omitempty" tf:"zonal,omitempty"`
}

// ClusterV1Spec defines the desired state of ClusterV1
type ClusterV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterV1Parameters `json:"forProvider"`
}

// ClusterV1Status defines the observed state of ClusterV1.
type ClusterV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterV1 is the Schema for the ClusterV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,selecteljet}
type ClusterV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterV1Spec   `json:"spec"`
	Status            ClusterV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterV1List contains a list of ClusterV1s
type ClusterV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterV1 `json:"items"`
}

// Repository type metadata.
var (
	ClusterV1_Kind             = "ClusterV1"
	ClusterV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterV1_Kind}.String()
	ClusterV1_KindAPIVersion   = ClusterV1_Kind + "." + CRDGroupVersion.String()
	ClusterV1_GroupVersionKind = CRDGroupVersion.WithKind(ClusterV1_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterV1{}, &ClusterV1List{})
}
