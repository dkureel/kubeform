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

type PostgresqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlFirewallRuleSpec   `json:"spec,omitempty"`
	Status            PostgresqlFirewallRuleStatus `json:"status,omitempty"`
}

type PostgresqlFirewallRuleSpec struct {
	EndIPAddress      string                    `json:"endIPAddress" tf:"end_ip_address"`
	Name              string                    `json:"name" tf:"name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName        string                    `json:"serverName" tf:"server_name"`
	StartIPAddress    string                    `json:"startIPAddress" tf:"start_ip_address"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PostgresqlFirewallRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PostgresqlFirewallRuleList is a list of PostgresqlFirewallRules
type PostgresqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PostgresqlFirewallRule CRD objects
	Items []PostgresqlFirewallRule `json:"items,omitempty"`
}
