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

type LicenseV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	Servers []LicenseV2ServersObservation `json:"servers,omitempty" tf:"servers,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type LicenseV2Parameters struct {

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LicenseV2ServersObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type LicenseV2ServersParameters struct {
}

// LicenseV2Spec defines the desired state of LicenseV2
type LicenseV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LicenseV2Parameters `json:"forProvider"`
}

// LicenseV2Status defines the observed state of LicenseV2.
type LicenseV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LicenseV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseV2 is the Schema for the LicenseV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,selecteljet}
type LicenseV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicenseV2Spec   `json:"spec"`
	Status            LicenseV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseV2List contains a list of LicenseV2s
type LicenseV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseV2 `json:"items"`
}

// Repository type metadata.
var (
	LicenseV2_Kind             = "LicenseV2"
	LicenseV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LicenseV2_Kind}.String()
	LicenseV2_KindAPIVersion   = LicenseV2_Kind + "." + CRDGroupVersion.String()
	LicenseV2_GroupVersionKind = CRDGroupVersion.WithKind(LicenseV2_Kind)
)

func init() {
	SchemeBuilder.Register(&LicenseV2{}, &LicenseV2List{})
}
