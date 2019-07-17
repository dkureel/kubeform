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

type HdinsightKafkaCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightKafkaClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightKafkaClusterStatus `json:"status,omitempty"`
}

type HdinsightKafkaClusterSpecComponentVersion struct {
	Kafka string `json:"kafka" tf:"kafka"`
}

type HdinsightKafkaClusterSpecGateway struct {
	Enabled  bool   `json:"enabled" tf:"enabled"`
	Password string `json:"password" tf:"password"`
	Username string `json:"username" tf:"username"`
}

type HdinsightKafkaClusterSpecRolesHeadNode struct {
	// +optional
	Password string `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightKafkaClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount     int `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
	NumberOfDisksPerNode int `json:"numberOfDisksPerNode" tf:"number_of_disks_per_node"`
	// +optional
	Password string `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID            string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	TargetInstanceCount int    `json:"targetInstanceCount" tf:"target_instance_count"`
	Username            string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightKafkaClusterSpecRolesZookeeperNode struct {
	// +optional
	Password string `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightKafkaClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightKafkaClusterSpecRolesHeadNode `json:"headNode" tf:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightKafkaClusterSpecRolesWorkerNode `json:"workerNode" tf:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightKafkaClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightKafkaClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"isDefault" tf:"is_default"`
	StorageAccountKey  string `json:"storageAccountKey" tf:"storage_account_key"`
	StorageContainerID string `json:"storageContainerID" tf:"storage_container_id"`
}

type HdinsightKafkaClusterSpec struct {
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightKafkaClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightKafkaClusterSpecGateway `json:"gateway" tf:"gateway"`
	Location          string                             `json:"location" tf:"location"`
	Name              string                             `json:"name" tf:"name"`
	ResourceGroupName string                             `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightKafkaClusterSpecRoles          `json:"roles" tf:"roles"`
	StorageAccount []HdinsightKafkaClusterSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
	Tier           string                                    `json:"tier" tf:"tier"`
	ProviderRef    core.LocalObjectReference                 `json:"providerRef" tf:"-"`
}

type HdinsightKafkaClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightKafkaClusterList is a list of HdinsightKafkaClusters
type HdinsightKafkaClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightKafkaCluster CRD objects
	Items []HdinsightKafkaCluster `json:"items,omitempty"`
}
