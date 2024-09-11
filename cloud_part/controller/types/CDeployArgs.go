package types

import (
	corev1 "k8s.io/api/core/v1"
)

// Depoloyment information of k8s

// k8s pod arguments
type CPodArgs struct {
	// About the image of container
	// type [image]:[version]
	Image string
	// 标签选择器，当使用服务时，k8s会自动选择一个pod
	Label map[string]string
	// 容器的名称，其他pod可以使用这个来通讯
	Name string
	// 公开的端口，部署一个服务来公开这个端口
	ExposePorts []ExposePort
	// 容器的入口点。这相当于Dockerfile中的ENTRYPOINT变量
	Entrypoint []string
	// 容器的命令。这相当于Dockerfile中的CMD变量
	Command []string
	/*
		Environment variables
		Each entry is a key value string pair.
		Example:
		- Name: "KAFKA_BROKER_ID",
		  Value: "0",
	*/
	Environment []Environment
	// 容器的内存限制
	// 存在两种限制，硬限制和软限制
	MemLimit     int64
	MemsoftLimit int64
	// cpu限制,pod可以使用的CPU数量，可以是小数
	CPULimit float64
	// 在运行时保证使用的CPU数量
	CPUSoftLimit float64
	Volumes      []CVolume
	//加速器的名称。必须手动定义具有GPU支持的Kubernetes节点，使用' kubectl label nodes <node-with-certain-gpu> accelerator=<certain-gpu-name>
	Accelerator string
	Nvidia      int64
	AMD         int64
	//镜像拉取策略。控制是否从 Docker Hub 或其他镜像仓库拉取镜像，或者仅使用本地缓存。
	PullPolicy corev1.PullPolicy
}

// 容器的端口格式
type ExposePort struct {
	HostPort      uint16
	ContainerPort uint16
	Protocol      Protocol
}

type Environment struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// 在容器中，将容器的目录挂载到主机目录
type CVolume struct {
	MountPoint string
	ReadOnly   bool
	Type       VolumeType
	HostPath   *HostPath
	//网络文件系统（Network File System）卷。这种类型的卷允许 Pod 访问网络上的文件系统
	NFS *NFS
	//Amazon Web Services Elastic Block Store (EBS) 卷。这种类型的卷允许 Pod 使用 AWS 提供的 EBS 卷。
	AwsEbs *AwsEbs
	//空目录卷。这种类型的卷在 Pod 启动时创建，在 Pod 终止时销毁。
	EmptyDir *EmptyDir
	//ConfigMap 卷。这种类型的卷允许 Pod 访问 Kubernetes 中的 ConfigMap 对象。
	ConfigMap *ConfigMapVolume
}

type HostPath struct {
	Path      string
	MountType corev1.HostPathType
	//表示挂载类型。corev1.HostPathType 是一个枚举类型，用于指定如何解释 Path 字段中的路径
	NodeHostname string
	//可以通过运行 kubectl get nodes 命令来获取节点的主机名列表
}

// The pod will mount to a path in an NFS server as a volume
type NFS struct {
	Server string
	Path   string
}

// AWS EBS volume
// You must first create an EBS volume on AWS
// This mounts the entire EBS disk to the target pod. You can mitigate this by partitioning the disk and provision it to multiple pods
type AwsEbs struct {
	//Volume ID of the EBS volume
	VolumeID string
	//File system type. Such as ext4
	FSType    string
	Partition int32
}

// Empty directory
// Creates an empty directory as a volume. This volume persists until the pod is removed or if the cluster is shut down
// Can be useful for pods to store data and persist through crashes
type EmptyDir struct {
	//Where to store this emptyDir
	//Default goes to disk
	//Memory goes to tmpfs (In memory, be careful of memory usage. Gets cleared when server shuts down)
	//Huge pages (In virtual memory. No size limits. Gets cleared when server shuts down)
	StorageMedium corev1.StorageMedium
	//Size limit in bytes - nil for unlimited
	SizeLimit *int64
}

// ConfigMap volume
// A ConfigMap is a key-value map that is stored in Kubernetes
// In most cases, this is a static file such as configurations files that you want to be available to your pods
// To mount a ConfigMap, you must first create a ConfigMap in Kubernetes
type ConfigMapVolume struct {
	//Name of the ConfigMap
	Name string
	//Permissions to apply to the volume
	Mode *int32
}

type VolumeType string

const (
	TypeHostPath  VolumeType = "hostpath"
	TypeNFS       VolumeType = "nfs"
	TypeAWSEBS    VolumeType = "awsebs"
	TypeEmptyDir  VolumeType = "emptyDir"
	TypeConfigMap VolumeType = "configMap"
)

type Protocol string

const (
	TCP  Protocol = "TCP"
	UDP  Protocol = "UDP"
	SCTP Protocol = "SCTP"
)
