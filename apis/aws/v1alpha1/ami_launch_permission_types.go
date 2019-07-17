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

type AmiLaunchPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiLaunchPermissionSpec   `json:"spec,omitempty"`
	Status            AmiLaunchPermissionStatus `json:"status,omitempty"`
}

type AmiLaunchPermissionSpec struct {
	AccountID   string                    `json:"accountID" tf:"account_id"`
	ImageID     string                    `json:"imageID" tf:"image_id"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AmiLaunchPermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AmiLaunchPermissionList is a list of AmiLaunchPermissions
type AmiLaunchPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AmiLaunchPermission CRD objects
	Items []AmiLaunchPermission `json:"items,omitempty"`
}
