package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DxHostedPublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPublicVirtualInterfaceSpec   `json:"spec,omitempty"`
	Status            DxHostedPublicVirtualInterfaceStatus `json:"status,omitempty"`
}

type DxHostedPublicVirtualInterfaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressFamily string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AwsDevice string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`
	BgpAsn    int    `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey   string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`
	ConnectionID string `json:"connectionID" tf:"connection_id"`
	// +optional
	CustomerAddress string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`
	Name            string `json:"name" tf:"name"`
	OwnerAccountID  string `json:"ownerAccountID" tf:"owner_account_id"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	RouteFilterPrefixes []string `json:"routeFilterPrefixes" tf:"route_filter_prefixes"`
	Vlan                int      `json:"vlan" tf:"vlan"`
}



type DxHostedPublicVirtualInterfaceStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *DxHostedPublicVirtualInterfaceSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxHostedPublicVirtualInterfaceList is a list of DxHostedPublicVirtualInterfaces
type DxHostedPublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxHostedPublicVirtualInterface CRD objects
	Items []DxHostedPublicVirtualInterface `json:"items,omitempty"`
}