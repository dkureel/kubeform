package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type NetworkInterfaceApplicationSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceApplicationSecurityGroupAssociationSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceApplicationSecurityGroupAssociationStatus `json:"status,omitempty"`
}

type NetworkInterfaceApplicationSecurityGroupAssociationSpec struct {
	ApplicationSecurityGroupID string                    `json:"applicationSecurityGroupID" tf:"application_security_group_id"`
	IpConfigurationName        string                    `json:"ipConfigurationName" tf:"ip_configuration_name"`
	NetworkInterfaceID         string                    `json:"networkInterfaceID" tf:"network_interface_id"`
	ProviderRef                core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NetworkInterfaceApplicationSecurityGroupAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceApplicationSecurityGroupAssociationList is a list of NetworkInterfaceApplicationSecurityGroupAssociations
type NetworkInterfaceApplicationSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterfaceApplicationSecurityGroupAssociation CRD objects
	Items []NetworkInterfaceApplicationSecurityGroupAssociation `json:"items,omitempty"`
}
