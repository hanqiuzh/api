package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Cluster"

// ClusterManagementAddOn represents the registration of an add-on to the cluster manager.
// This resource allows the user to discover which add-on is available for the cluster manager and
// also provides metadata information about the add-on.
// The resource also provides a linkage to ManagedClusterAddOn, the name of the ClusterManagementAddOn
// resource will be use for the namespace scoped ManagedClusterAddOn resource
// ClusterManagementAddOn is a cluster scoped resource.
type ClusterManagementAddOn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents a desired configuration for the agent on the cluster manager add-on
	Spec ClusterManagementAddOnSpec `json:"spec"`

	// Status represents the current status of cluster manager add-on
	// +optional
	Status ClusterManagementAddOnStatus `json:"status,omitempty"`
}

// ClusterManagementAddOnSpec provides the information of add-on CustomResourceDefinition.
type ClusterManagementAddOnSpec struct {
	// DisplayName represents the name that will be displayed.
	// +required
	DisplayName string `json:"displayName"`

	// Description represents the detailed description of the add-on.
	// +optional
	Description string `json:"description"`

	// AddOnConfigCRD is a reference to the name of the CRD that driving the configuration of the add-on
	// Note: there may be a case where a single add-on config CRD controls multiple related add-ons,
	// in this case multiple ClusterManagementAddOn resource should be created.
	// +required
	AddOnConfigCRD string `json:"addOnConfigCRD"`
}

// ClusterManagementAddOnStatus represents the current status of cluster manager add-on.
type ClusterManagementAddOnStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ClusterManagementAddOnList is a collection of cluster manager add-ons.
type ClusterManagementAddOnList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of cluster manager add-ons.
	Items []ClusterManagementAddOn `json:"items"`
}
