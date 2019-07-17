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

type ComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

type ComputeInstanceGroupManagerSpecAutoHealingPolicies struct {
	HealthCheck     string `json:"healthCheck" tf:"health_check"`
	InitialDelaySec int    `json:"initialDelaySec" tf:"initial_delay_sec"`
}

type ComputeInstanceGroupManagerSpecNamedPort struct {
	Name string `json:"name" tf:"name"`
	Port int    `json:"port" tf:"port"`
}

type ComputeInstanceGroupManagerSpecRollingUpdatePolicy struct {
	// +optional
	MaxSurgeFixed int `json:"maxSurgeFixed,omitempty" tf:"max_surge_fixed,omitempty"`
	// +optional
	MaxSurgePercent int `json:"maxSurgePercent,omitempty" tf:"max_surge_percent,omitempty"`
	// +optional
	MaxUnavailableFixed int `json:"maxUnavailableFixed,omitempty" tf:"max_unavailable_fixed,omitempty"`
	// +optional
	MaxUnavailablePercent int `json:"maxUnavailablePercent,omitempty" tf:"max_unavailable_percent,omitempty"`
	// +optional
	MinReadySec   int    `json:"minReadySec,omitempty" tf:"min_ready_sec,omitempty"`
	MinimalAction string `json:"minimalAction" tf:"minimal_action"`
	Type          string `json:"type" tf:"type"`
}

type ComputeInstanceGroupManagerSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	AutoHealingPolicies []ComputeInstanceGroupManagerSpecAutoHealingPolicies `json:"autoHealingPolicies,omitempty" tf:"auto_healing_policies,omitempty"`
	BaseInstanceName    string                                               `json:"baseInstanceName" tf:"base_instance_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	InstanceTemplate string `json:"instanceTemplate,omitempty" tf:"instance_template,omitempty"`
	Name             string `json:"name" tf:"name"`
	// +optional
	NamedPort []ComputeInstanceGroupManagerSpecNamedPort `json:"namedPort,omitempty" tf:"named_port,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	RollingUpdatePolicy []ComputeInstanceGroupManagerSpecRollingUpdatePolicy `json:"rollingUpdatePolicy,omitempty" tf:"rolling_update_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TargetPools []string `json:"targetPools,omitempty" tf:"target_pools,omitempty"`
	// +optional
	UpdateStrategy string `json:"updateStrategy,omitempty" tf:"update_strategy,omitempty"`
	// +optional
	WaitForInstances bool                      `json:"waitForInstances,omitempty" tf:"wait_for_instances,omitempty"`
	ProviderRef      core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeInstanceGroupManagerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceGroupManagerList is a list of ComputeInstanceGroupManagers
type ComputeInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstanceGroupManager CRD objects
	Items []ComputeInstanceGroupManager `json:"items,omitempty"`
}
