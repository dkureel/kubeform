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

type PubsubTopicIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PubsubTopicIamPolicySpec   `json:"spec,omitempty"`
	Status            PubsubTopicIamPolicyStatus `json:"status,omitempty"`
}

type PubsubTopicIamPolicySpec struct {
	PolicyData  string                    `json:"policyData" tf:"policy_data"`
	Topic       string                    `json:"topic" tf:"topic"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PubsubTopicIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PubsubTopicIamPolicyList is a list of PubsubTopicIamPolicys
type PubsubTopicIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PubsubTopicIamPolicy CRD objects
	Items []PubsubTopicIamPolicy `json:"items,omitempty"`
}
