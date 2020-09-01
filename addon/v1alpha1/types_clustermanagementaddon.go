package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Cluster"
// +kubebuilder:printcolumn:name="DISPLAY NAME",type=string,JSONPath=`.spec.displayName`
// +kubebuilder:printcolumn:name="CRD NAME",type=string,JSONPath=`.spec.addOnConfigCRD`

// ClusterManagementAddOn represents the registration of an add-on to the cluster management.
// This resource allows the user to discover which add-on is available for the cluster management and
// also provides metadata information about the add-on.
// The resource also provides a linkage to ManagedClusterAddOn, the name of the ClusterManagementAddOn
// resource will be used for the namespace-scoped ManagedClusterAddOn resource
// ClusterManagementAddOn is a cluster-scoped resource.
type ClusterManagementAddOn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec represents a desired configuration for the agent on the cluster management add-on.
	Spec ClusterManagementAddOnSpec `json:"spec"`

	// status represents the current status of cluster management add-on.
	// +optional
	Status ClusterManagementAddOnStatus `json:"status,omitempty"`
}

// ClusterManagementAddOnSpec provides the information of add-on CustomResourceDefinition.
type ClusterManagementAddOnSpec struct {
	// displayName represents the name that will be displayed.
	// +required
	DisplayName string `json:"displayName"`

	// description represents the detailed description of the add-on.
	// +optional
	Description string `json:"description"`

	// addOnConfigCRD is a reference to the name of the CRD that configures the add-on.
	// In scenario where a multiple add-ons share the same add-on config CRD multiple ClusterManagementAddOn resource need to be created.
	// +required
	AddOnConfigCRD string `json:"addOnConfigCRD"`
}

// ClusterManagementAddOnStatus represents the current status of cluster management add-on.
type ClusterManagementAddOnStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ClusterManagementAddOnList is a collection of cluster management add-ons.
type ClusterManagementAddOnList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of cluster management add-ons.
	Items []ClusterManagementAddOn `json:"items"`
}
