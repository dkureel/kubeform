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

type FloatingIP struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FloatingIPSpec   `json:"spec,omitempty"`
	Status            FloatingIPStatus `json:"status,omitempty"`
}

type FloatingIPSpec struct {
	// +optional
	DropletID   int                       `json:"dropletID,omitempty" tf:"droplet_id,omitempty"`
	Region      string                    `json:"region" tf:"region"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FloatingIPStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FloatingIPList is a list of FloatingIPs
type FloatingIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FloatingIP CRD objects
	Items []FloatingIP `json:"items,omitempty"`
}
