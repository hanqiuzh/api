package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Cluster"

// ClusterManagerAddOn represents the registration of an addon
// It allows the user to discover which addon is available for the cluster manager.
// It also provides metadata information about the addon
// ClusterManagerAddOn is a cluster scoped resource.
type ClusterManagerAddOn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents a desired configuration for the agent on the cluster manager addon
	Spec ClusterManagerAddOnSpec `json:"spec"`

	// Status represents the current status of cluster manager addon
	// +optional
	Status ClusterManagerAddOnStatus `json:"status,omitempty"`
}

// ClusterManagerAddOnSpec provides the information of addon CRD, related ManagedClusterAddon
// names, and addon information we will use for UI
type ClusterManagerAddOnSpec struct {
	// DisplayName represents the name that will be displayed.
	// +required
	DisplayName string `json:"displayName"`

	// Description represents the detail description of the addon.
	// +required
	Description string `json:"description"`

	// AddOnConfigCRD is the name for the CRD name of the addon CRs, which is used for addon configurations.
	// It will inform the user which resource will the add-on controller/operator use to control the addon.
	// +required
	AddOnConfigCRD string `json:"addOnConfigCRD"`

	// ManagedClusterAddOnName is the metadata.name of the referenced ManagedClusterAddon CR.
	// The name will be use to associate the ManagedClusterAddOn resource to ClusterManagerAddon.
	// +required
	ManagedClusterAddOnName string `json:"name"`
}

// ClusterManagerAddOnStatus represents the current status of cluster manager addon.
type ClusterManagerAddOnStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ClusterManagerAddOnList is a collection of cluster manager addons.
type ClusterManagerAddOnList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of cluster manager addons.
	Items []ClusterManagerAddOn `json:"items"`
}
