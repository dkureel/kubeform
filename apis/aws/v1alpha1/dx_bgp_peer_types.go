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

type DxBGPPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxBGPPeerSpec   `json:"spec,omitempty"`
	Status            DxBGPPeerStatus `json:"status,omitempty"`
}

type DxBGPPeerSpec struct {
	AddressFamily      string                    `json:"addressFamily" tf:"address_family"`
	BgpAsn             int                       `json:"bgpAsn" tf:"bgp_asn"`
	VirtualInterfaceID string                    `json:"virtualInterfaceID" tf:"virtual_interface_id"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DxBGPPeerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxBGPPeerList is a list of DxBGPPeers
type DxBGPPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxBGPPeer CRD objects
	Items []DxBGPPeer `json:"items,omitempty"`
}
