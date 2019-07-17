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

type ComputeRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterSpec   `json:"spec,omitempty"`
	Status            ComputeRouterStatus `json:"status,omitempty"`
}

type ComputeRouterSpecBgpAdvertisedIPRanges struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Range string `json:"range,omitempty" tf:"range,omitempty"`
}

type ComputeRouterSpecBgp struct {
	// +optional
	AdvertiseMode string `json:"advertiseMode,omitempty" tf:"advertise_mode,omitempty"`
	// +optional
	AdvertisedGroups []string `json:"advertisedGroups,omitempty" tf:"advertised_groups,omitempty"`
	// +optional
	AdvertisedIPRanges []ComputeRouterSpecBgpAdvertisedIPRanges `json:"advertisedIPRanges,omitempty" tf:"advertised_ip_ranges,omitempty"`
	Asn                int                                      `json:"asn" tf:"asn"`
}

type ComputeRouterSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Bgp []ComputeRouterSpecBgp `json:"bgp,omitempty" tf:"bgp,omitempty"`
	// +optional
	Description string                    `json:"description,omitempty" tf:"description,omitempty"`
	Name        string                    `json:"name" tf:"name"`
	Network     string                    `json:"network" tf:"network"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeRouterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterList is a list of ComputeRouters
type ComputeRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouter CRD objects
	Items []ComputeRouter `json:"items,omitempty"`
}
