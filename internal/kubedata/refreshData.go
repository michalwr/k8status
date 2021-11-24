package kubedata


func  (Kube *KubernetesCluster) RefreshData()() {
    	tmpNodeStatus:=GetNodeStatus(Kube.Clientset)
		tmpPodStatus:=GetpodsProps(Kube.Clientset)
		Kube.NodeProps=tmpNodeStatus
		Kube.PodProps=tmpPodStatus
}