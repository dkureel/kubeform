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

type ComputeAddress struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeAddressSpec   `json:"spec,omitempty"`
	Status            ComputeAddressStatus `json:"status,omitempty"`
}

type ComputeAddressSpec struct {
	// +optional
	AddressType string `json:"addressType,omitempty" tf:"address_type,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// Deprecated
	Labels      map[string]string         `json:"labels,omitempty" tf:"labels,omitempty"`
	Name        string                    `json:"name" tf:"name"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeAddressStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeAddressList is a list of ComputeAddresss
type ComputeAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeAddress CRD objects
	Items []ComputeAddress `json:"items,omitempty"`
}
