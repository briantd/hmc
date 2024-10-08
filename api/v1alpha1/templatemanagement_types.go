// Copyright 2024
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=tm,scope=Cluster

// TemplateManagement is the Schema for the template management API
type TemplateManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TemplateManagementSpec   `json:"spec,omitempty"`
	Status TemplateManagementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TemplateManagementList contains a list of TemplateManagement
type TemplateManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TemplateManagement `json:"items"`
}

// TemplateManagementSpec defines the desired state of TemplateManagement
type TemplateManagementSpec struct {
	// AccessRules is the list of access rules. Each AccessRule enforces
	// Templates distribution to the TargetNamespaces.
	AccessRules []AccessRule `json:"accessRules,omitempty"`
}

// AccessRule is the definition of the TemplateManagement access rule. Each AccessRule enforces
// Templates distribution to the TargetNamespaces
type AccessRule struct {
	// TargetNamespaces defines the namespaces where selected Templates will be distributed.
	// Templates will be distributed to all namespaces if unset.
	// +optional
	TargetNamespaces TargetNamespaces `json:"targetNamespaces,omitempty"`
	// ClusterTemplateChains lists the names of ClusterTemplateChains whose ClusterTemplates
	// will be distributed to all namespaces specified in TargetNamespaces.
	// +optional
	ClusterTemplateChains []string `json:"clusterTemplateChains,omitempty"`
	// ServiceTemplateChains lists the names of ServiceTemplateChains whose ServiceTemplates
	// will be distributed to all namespaces specified in TargetNamespaces.
	// +optional
	ServiceTemplateChains []string `json:"serviceTemplateChains,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="((has(self.stringSelector) ? 1 : 0) + (has(self.selector) ? 1 : 0) + (has(self.list) ? 1 : 0)) <= 1", message="only one of spec.targetNamespaces.selector or spec.targetNamespaces.stringSelector spec.targetNamespaces.list can be specified"

// TargetNamespaces defines the list of namespaces or the label selector to select namespaces
type TargetNamespaces struct {
	// StringSelector is a label query to select namespaces.
	// Mutually exclusive with Selector.
	// +optional
	StringSelector string `json:"stringSelector,omitempty"`
	// Selector is a structured label query to select namespaces.
	// Mutually exclusive with StringSelector.
	// +optional
	Selector metav1.LabelSelector `json:"selector,omitempty"`
	// List is the list of namespaces to select.
	// +optional
	List []string `json:"list,omitempty"`
}

// TemplateManagementStatus defines the observed state of TemplateManagement
type TemplateManagementStatus struct {
	// ObservedGeneration is the last observed generation.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// Current reflects the applied access rules configuration.
	Current []AccessRule `json:"current,omitempty"`
}

func init() {
	SchemeBuilder.Register(&TemplateManagement{}, &TemplateManagementList{})
}
