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

type BatchJobDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchJobDefinitionSpec   `json:"spec,omitempty"`
	Status            BatchJobDefinitionStatus `json:"status,omitempty"`
}

type BatchJobDefinitionSpecRetryStrategy struct {
	// +optional
	Attempts int `json:"attempts,omitempty" tf:"attempts,omitempty"`
}

type BatchJobDefinitionSpecTimeout struct {
	// +optional
	AttemptDurationSeconds int `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds,omitempty"`
}

type BatchJobDefinitionSpec struct {
	// +optional
	ContainerProperties string `json:"containerProperties,omitempty" tf:"container_properties,omitempty"`
	Name                string `json:"name" tf:"name"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetryStrategy []BatchJobDefinitionSpecRetryStrategy `json:"retryStrategy,omitempty" tf:"retry_strategy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Timeout     []BatchJobDefinitionSpecTimeout `json:"timeout,omitempty" tf:"timeout,omitempty"`
	Type        string                          `json:"type" tf:"type"`
	ProviderRef core.LocalObjectReference       `json:"providerRef" tf:"-"`
}

type BatchJobDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchJobDefinitionList is a list of BatchJobDefinitions
type BatchJobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchJobDefinition CRD objects
	Items []BatchJobDefinition `json:"items,omitempty"`
}
