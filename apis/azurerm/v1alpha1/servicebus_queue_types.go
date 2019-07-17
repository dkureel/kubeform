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

type ServicebusQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusQueueSpec   `json:"spec,omitempty"`
	Status            ServicebusQueueStatus `json:"status,omitempty"`
}

type ServicebusQueueSpec struct {
	// +optional
	DeadLetteringOnMessageExpiration bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`
	// +optional
	// Deprecated
	EnableBatchedOperations bool `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`
	// +optional
	EnableExpress bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`
	// +optional
	EnablePartitioning bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	MaxDeliveryCount int    `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
	Name             string `json:"name" tf:"name"`
	NamespaceName    string `json:"namespaceName" tf:"namespace_name"`
	// +optional
	RequiresDuplicateDetection bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`
	// +optional
	RequiresSession   bool   `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// Deprecated
	SupportOrdering bool                      `json:"supportOrdering,omitempty" tf:"support_ordering,omitempty"`
	ProviderRef     core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ServicebusQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusQueueList is a list of ServicebusQueues
type ServicebusQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusQueue CRD objects
	Items []ServicebusQueue `json:"items,omitempty"`
}
