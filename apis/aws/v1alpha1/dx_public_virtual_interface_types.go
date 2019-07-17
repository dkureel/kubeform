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

type DxPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxPublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            DxPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type DxPublicVirtualInterfaceSpec struct {
	AddressFamily string `json:"addressFamily" tf:"address_family"`
	BgpAsn        int    `json:"bgpAsn" tf:"bgp_asn"`
	ConnectionID  string `json:"connectionID" tf:"connection_id"`
	Name          string `json:"name" tf:"name"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	RouteFilterPrefixes []string `json:"routeFilterPrefixes" tf:"route_filter_prefixes"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	Vlan        int                       `json:"vlan" tf:"vlan"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DxPublicVirtualInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxPublicVirtualInterfaceList is a list of DxPublicVirtualInterfaces
type DxPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxPublicVirtualInterface CRD objects
	Items []DxPublicVirtualInterface `json:"items,omitempty"`
}
