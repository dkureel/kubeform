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

type StorageTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageTableSpec   `json:"spec,omitempty"`
	Status            StorageTableStatus `json:"status,omitempty"`
}

type StorageTableSpecAclAccessPolicy struct {
	Expiry      string `json:"expiry" tf:"expiry"`
	Permissions string `json:"permissions" tf:"permissions"`
	Start       string `json:"start" tf:"start"`
}

type StorageTableSpecAcl struct {
	// +optional
	AccessPolicy []StorageTableSpecAclAccessPolicy `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`
	ID           string                            `json:"ID" tf:"id"`
}

type StorageTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Acl  []StorageTableSpecAcl `json:"acl,omitempty" tf:"acl,omitempty"`
	Name string                `json:"name" tf:"name"`
	// +optional
	// Deprecated
	ResourceGroupName  string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
}



type StorageTableStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *StorageTableSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageTableList is a list of StorageTables
type StorageTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageTable CRD objects
	Items []StorageTable `json:"items,omitempty"`
}