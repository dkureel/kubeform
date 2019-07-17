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

type KeyVaultCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultCertificateSpec   `json:"spec,omitempty"`
	Status            KeyVaultCertificateStatus `json:"status,omitempty"`
}

type KeyVaultCertificateSpecCertificate struct {
	Contents string `json:"contents" tf:"contents"`
	// +optional
	Password string `json:"password,omitempty" tf:"password,omitempty"`
}

type KeyVaultCertificateSpecCertificatePolicyIssuerParameters struct {
	Name string `json:"name" tf:"name"`
}

type KeyVaultCertificateSpecCertificatePolicyKeyProperties struct {
	Exportable bool   `json:"exportable" tf:"exportable"`
	KeySize    int    `json:"keySize" tf:"key_size"`
	KeyType    string `json:"keyType" tf:"key_type"`
	ReuseKey   bool   `json:"reuseKey" tf:"reuse_key"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction struct {
	ActionType string `json:"actionType" tf:"action_type"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger struct {
	// +optional
	DaysBeforeExpiry int `json:"daysBeforeExpiry,omitempty" tf:"days_before_expiry,omitempty"`
	// +optional
	LifetimePercentage int `json:"lifetimePercentage,omitempty" tf:"lifetime_percentage,omitempty"`
}

type KeyVaultCertificateSpecCertificatePolicyLifetimeAction struct {
	// +kubebuilder:validation:MaxItems=1
	Action []KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction `json:"action" tf:"action"`
	// +kubebuilder:validation:MaxItems=1
	Trigger []KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger `json:"trigger" tf:"trigger"`
}

type KeyVaultCertificateSpecCertificatePolicySecretProperties struct {
	ContentType string `json:"contentType" tf:"content_type"`
}

type KeyVaultCertificateSpecCertificatePolicy struct {
	// +kubebuilder:validation:MaxItems=1
	IssuerParameters []KeyVaultCertificateSpecCertificatePolicyIssuerParameters `json:"issuerParameters" tf:"issuer_parameters"`
	// +kubebuilder:validation:MaxItems=1
	KeyProperties []KeyVaultCertificateSpecCertificatePolicyKeyProperties `json:"keyProperties" tf:"key_properties"`
	// +optional
	LifetimeAction []KeyVaultCertificateSpecCertificatePolicyLifetimeAction `json:"lifetimeAction,omitempty" tf:"lifetime_action,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	SecretProperties []KeyVaultCertificateSpecCertificatePolicySecretProperties `json:"secretProperties" tf:"secret_properties"`
}

type KeyVaultCertificateSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Certificate []KeyVaultCertificateSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	CertificatePolicy []KeyVaultCertificateSpecCertificatePolicy `json:"certificatePolicy" tf:"certificate_policy"`
	Name              string                                     `json:"name" tf:"name"`
	ProviderRef       core.LocalObjectReference                  `json:"providerRef" tf:"-"`
}

type KeyVaultCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultCertificateList is a list of KeyVaultCertificates
type KeyVaultCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVaultCertificate CRD objects
	Items []KeyVaultCertificate `json:"items,omitempty"`
}
