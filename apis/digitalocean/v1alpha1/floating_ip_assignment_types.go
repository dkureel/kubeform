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

type FloatingIPAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FloatingIPAssignmentSpec   `json:"spec,omitempty"`
	Status            FloatingIPAssignmentStatus `json:"status,omitempty"`
}

type FloatingIPAssignmentSpec struct {
	DropletID   int                       `json:"dropletID" tf:"droplet_id"`
	IpAddress   string                    `json:"ipAddress" tf:"ip_address"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FloatingIPAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FloatingIPAssignmentList is a list of FloatingIPAssignments
type FloatingIPAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FloatingIPAssignment CRD objects
	Items []FloatingIPAssignment `json:"items,omitempty"`
}
