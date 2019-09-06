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

type StorageTableEntity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageTableEntitySpec   `json:"spec,omitempty"`
	Status            StorageTableEntityStatus `json:"status,omitempty"`
}

type StorageTableEntitySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Entity             map[string]string `json:"entity" tf:"entity"`
	PartitionKey       string            `json:"partitionKey" tf:"partition_key"`
	RowKey             string            `json:"rowKey" tf:"row_key"`
	StorageAccountName string            `json:"storageAccountName" tf:"storage_account_name"`
	TableName          string            `json:"tableName" tf:"table_name"`
}



type StorageTableEntityStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *StorageTableEntitySpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageTableEntityList is a list of StorageTableEntitys
type StorageTableEntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageTableEntity CRD objects
	Items []StorageTableEntity `json:"items,omitempty"`
}