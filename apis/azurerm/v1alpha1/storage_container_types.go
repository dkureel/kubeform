package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type StorageContainer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageContainerSpec   `json:"spec,omitempty"`
	Status            StorageContainerStatus `json:"status,omitempty"`
}

type StorageContainerSpec struct {
	// +optional
	ContainerAccessType string `json:"container_access_type,omitempty"`
	Name                string `json:"name"`
	ResourceGroupName   string `json:"resource_group_name"`
	StorageAccountName  string `json:"storage_account_name"`
}

type StorageContainerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageContainerList is a list of StorageContainers
type StorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageContainer CRD objects
	Items []StorageContainer `json:"items,omitempty"`
}