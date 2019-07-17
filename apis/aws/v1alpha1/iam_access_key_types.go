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

type IamAccessKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamAccessKeySpec   `json:"spec,omitempty"`
	Status            IamAccessKeyStatus `json:"status,omitempty"`
}

type IamAccessKeySpec struct {
	// +optional
	PgpKey      string                    `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`
	User        string                    `json:"user" tf:"user"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IamAccessKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamAccessKeyList is a list of IamAccessKeys
type IamAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamAccessKey CRD objects
	Items []IamAccessKey `json:"items,omitempty"`
}
