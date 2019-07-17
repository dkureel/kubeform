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

type InspectorAssessmentTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InspectorAssessmentTargetSpec   `json:"spec,omitempty"`
	Status            InspectorAssessmentTargetStatus `json:"status,omitempty"`
}

type InspectorAssessmentTargetSpec struct {
	Name string `json:"name" tf:"name"`
	// +optional
	ResourceGroupArn string                    `json:"resourceGroupArn,omitempty" tf:"resource_group_arn,omitempty"`
	ProviderRef      core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InspectorAssessmentTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InspectorAssessmentTargetList is a list of InspectorAssessmentTargets
type InspectorAssessmentTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InspectorAssessmentTarget CRD objects
	Items []InspectorAssessmentTarget `json:"items,omitempty"`
}
