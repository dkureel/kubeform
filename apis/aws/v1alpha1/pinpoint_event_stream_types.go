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

type PinpointEventStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointEventStreamSpec   `json:"spec,omitempty"`
	Status            PinpointEventStreamStatus `json:"status,omitempty"`
}

type PinpointEventStreamSpec struct {
	ApplicationId        string `json:"application_id"`
	DestinationStreamArn string `json:"destination_stream_arn"`
	RoleArn              string `json:"role_arn"`
}

type PinpointEventStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointEventStreamList is a list of PinpointEventStreams
type PinpointEventStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointEventStream CRD objects
	Items []PinpointEventStream `json:"items,omitempty"`
}