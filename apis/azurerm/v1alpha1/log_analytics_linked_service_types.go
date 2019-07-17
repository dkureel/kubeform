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

type LogAnalyticsLinkedService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsLinkedServiceSpec   `json:"spec,omitempty"`
	Status            LogAnalyticsLinkedServiceStatus `json:"status,omitempty"`
}

type LogAnalyticsLinkedServiceSpec struct {
	// +optional
	LinkedServiceName string                    `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	WorkspaceName     string                    `json:"workspaceName" tf:"workspace_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LogAnalyticsLinkedServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogAnalyticsLinkedServiceList is a list of LogAnalyticsLinkedServices
type LogAnalyticsLinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogAnalyticsLinkedService CRD objects
	Items []LogAnalyticsLinkedService `json:"items,omitempty"`
}
