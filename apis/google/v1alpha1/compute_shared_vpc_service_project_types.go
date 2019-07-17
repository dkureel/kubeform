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

type ComputeSharedVpcServiceProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSharedVpcServiceProjectSpec   `json:"spec,omitempty"`
	Status            ComputeSharedVpcServiceProjectStatus `json:"status,omitempty"`
}

type ComputeSharedVpcServiceProjectSpec struct {
	HostProject    string                    `json:"hostProject" tf:"host_project"`
	ServiceProject string                    `json:"serviceProject" tf:"service_project"`
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeSharedVpcServiceProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSharedVpcServiceProjectList is a list of ComputeSharedVpcServiceProjects
type ComputeSharedVpcServiceProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSharedVpcServiceProject CRD objects
	Items []ComputeSharedVpcServiceProject `json:"items,omitempty"`
}
