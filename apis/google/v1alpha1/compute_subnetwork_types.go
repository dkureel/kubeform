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

type ComputeSubnetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSubnetworkSpec   `json:"spec,omitempty"`
	Status            ComputeSubnetworkStatus `json:"status,omitempty"`
}

type ComputeSubnetworkSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EnableFlowLogs bool   `json:"enableFlowLogs,omitempty" tf:"enable_flow_logs,omitempty"`
	IpCIDRRange    string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	Name           string `json:"name" tf:"name"`
	Network        string `json:"network" tf:"network"`
	// +optional
	PrivateIPGoogleAccess bool                      `json:"privateIPGoogleAccess,omitempty" tf:"private_ip_google_access,omitempty"`
	ProviderRef           core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeSubnetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSubnetworkList is a list of ComputeSubnetworks
type ComputeSubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSubnetwork CRD objects
	Items []ComputeSubnetwork `json:"items,omitempty"`
}
