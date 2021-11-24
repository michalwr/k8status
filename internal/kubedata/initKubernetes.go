package kubedata

import "k8s.io/client-go/kubernetes"

type KubernetesCluster struct {
	Name      string
	Clientset *kubernetes.Clientset
	NodeProps []NodeProps
	PodProps  []PodProps
}

func InitKubernetes() KubernetesCluster {
	var Kube KubernetesCluster
	Kube.Clientset=Kube.getConfig()
	Kube.NodeProps=GetNodeStatus(Kube.Clientset)
	Kube.PodProps=GetpodsProps(Kube.Clientset)
	return Kube
}
