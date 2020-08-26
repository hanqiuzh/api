package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Cluster"

// ClusterManagerAddOn represents the desired state and current status of one/many
// managed cluster addons. ClusterManagerAddOn is a cluster scoped resource.
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
	// CRDRef is the CRD name of the addon CRs, which used for addon configurations
	// +required
	CRDRef string `json:"crdRef"`
	// AddOns is the array of addon information helps the UI to find and display each addons
	// +required
	AddOns []AddOnInfo `json:"addons"`
}

// AddOnInfo contains the name of ManagedClusterAddon CR, display name we will use for UI, and
// description of an addon
type AddOnInfo struct {
	// Name is the same as the metadata.name of the referenced ManagedClusterAddon CR
	// +required
	Name string `json:"name"`
	// DisplayName is the addon name displayed on UI
	DisplayName string `json:"displayName"`
	// Description is the description of the addon displayed on UI
	Description string `json:"description"`
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
