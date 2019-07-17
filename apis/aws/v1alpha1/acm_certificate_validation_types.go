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

type AcmCertificateValidation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmCertificateValidationSpec   `json:"spec,omitempty"`
	Status            AcmCertificateValidationStatus `json:"status,omitempty"`
}

type AcmCertificateValidationSpec struct {
	CertificateArn string `json:"certificateArn" tf:"certificate_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ValidationRecordFqdns []string                  `json:"validationRecordFqdns,omitempty" tf:"validation_record_fqdns,omitempty"`
	ProviderRef           core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AcmCertificateValidationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AcmCertificateValidationList is a list of AcmCertificateValidations
type AcmCertificateValidationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AcmCertificateValidation CRD objects
	Items []AcmCertificateValidation `json:"items,omitempty"`
}
