package kubedata

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"context"
)
type NodeProps struct {
	Name string
	Status string
}
func GetNodeStatus()([]NodeProps) {
	var NodesPoperties[] NodeProps
	clientset:=getConfig()
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



