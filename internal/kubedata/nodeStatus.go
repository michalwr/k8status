package kubedata

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)
type NodeProps struct {
	Name string
	Status string
}
func  GetNodeStatus(clientset  *kubernetes.Clientset)([]NodeProps) {
	var NodesPoperties[] NodeProps
	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	for _, vNode := range nodes.Items {
		for _, vConditions := range vNode.Status.Conditions {
			if vConditions.Type == "Ready"   {
				NodesPoperties = append(NodesPoperties,  NodeProps{Name: vNode.Name, Status: string(vConditions.Status)})
			}
		}
	}
	return NodesPoperties
}



