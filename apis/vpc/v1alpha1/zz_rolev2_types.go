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

type RoleV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoleV2Parameters struct {

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Required
	UserID *string `json:"userId" tf:"user_id,omitempty"`
}

// RoleV2Spec defines the desired state of RoleV2
type RoleV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleV2Parameters `json:"forProvider"`
}

// RoleV2Status defines the observed state of RoleV2.
type RoleV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleV2 is the Schema for the RoleV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,selecteljet}
type RoleV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleV2Spec   `json:"spec"`
	Status            RoleV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleV2List contains a list of RoleV2s
type RoleV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleV2 `json:"items"`
}

// Repository type metadata.
var (
	RoleV2_Kind             = "RoleV2"
	RoleV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleV2_Kind}.String()
	RoleV2_KindAPIVersion   = RoleV2_Kind + "." + CRDGroupVersion.String()
	RoleV2_GroupVersionKind = CRDGroupVersion.WithKind(RoleV2_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleV2{}, &RoleV2List{})
}
