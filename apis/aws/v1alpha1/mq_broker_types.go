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

type MqBroker struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MqBrokerSpec   `json:"spec,omitempty"`
	Status            MqBrokerStatus `json:"status,omitempty"`
}

type MqBrokerSpecLogs struct {
	// +optional
	Audit bool `json:"audit,omitempty" tf:"audit,omitempty"`
	// +optional
	General bool `json:"general,omitempty" tf:"general,omitempty"`
}

type MqBrokerSpecUser struct {
	// +optional
	ConsoleAccess bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Groups   []string `json:"groups,omitempty" tf:"groups,omitempty"`
	Password string   `json:"password" tf:"password"`
	Username string   `json:"username" tf:"username"`
}

type MqBrokerSpec struct {
	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool   `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	BrokerName              string `json:"brokerName" tf:"broker_name"`
	// +optional
	DeploymentMode   string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`
	EngineType       string `json:"engineType" tf:"engine_type"`
	EngineVersion    string `json:"engineVersion" tf:"engine_version"`
	HostInstanceType string `json:"hostInstanceType" tf:"host_instance_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logs []MqBrokerSpecLogs `json:"logs,omitempty" tf:"logs,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups" tf:"security_groups"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	User        []MqBrokerSpecUser        `json:"user" tf:"user"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MqBrokerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MqBrokerList is a list of MqBrokers
type MqBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MqBroker CRD objects
	Items []MqBroker `json:"items,omitempty"`
}
