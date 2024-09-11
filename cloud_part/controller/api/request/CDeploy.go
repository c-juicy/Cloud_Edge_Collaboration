package request

import (
	"CLOUD_PART/controller/types"
)

/*
所有的部署API请求函数。

参见https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry
*/

// 在Kubernetes中启动部署
// Kubernetes将自动提取所有需要的images和依赖项

// 用于在k8s中创建deployment
func CRunDeployment(deployargs types.CDployArgs)
