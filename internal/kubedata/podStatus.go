package kubedata

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)
type PodProps struct {
	Name string
	Status string
	Namespace string
	ContainerStatuses []ContainerStatuses
}
type ContainerStatuses struct {
	Name string
	Status bool
	ExitCode int32
	RestartCount int

}
func  GetpodsProps(clientset *kubernetes.Clientset)([]PodProps) {
	var props[] PodProps
	var containterExitCode int32
	//var containerNotReadyReason string 
	pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	for _, vPod := range pods.Items {
		//fmt.Fprintf(w, "%d\t%s\t%s\t%s\t\n", i+1, d.Name, d.Status.Phase, d.Status.Reason)
		//for subi, c := range d.Status.ContainerStatuses {
			//fmt.Fprintf(w, "\t%d\t%s\t%t\t%d\t%d\t%s\t\n", subi+1, c.Name, c.Ready, c.LastTerminationState.Terminated.ExitCode, c.RestartCount, c.LastTerminationState.Terminated.FinishedAt)
			props = append(props,  PodProps{Name: vPod.Name, Status: string(vPod.Status.Phase), Namespace: vPod.Namespace})
			if vPod.Status.Phase == "Running" {
			for _, vContainer := range  vPod.Status.ContainerStatuses {
			containterExitCode=0
			if vContainer.RestartCount > 0 {
            containterExitCode=vContainer.LastTerminationState.Terminated.ExitCode
			}
			//fmt.Println(vPod.Name)
			//fmt.Println(vContainer.RestartCount)
			//fmt.Println(*vContainer.Started)
			//fmt.Println(vContainer.Ready)
			//fmt.Println(vContainer.State)
            // if !vContainer.Ready && !*vContainer.Started  {
			//     containerNotReadyReason=vContainer.State.Waiting.Reason
			// }
		        props[len(props)-1].ContainerStatuses = append(props[len(props)-1].ContainerStatuses,  ContainerStatuses{Name: vContainer.Name, Status: vContainer.Ready, ExitCode: containterExitCode, RestartCount: int(vContainer.RestartCount) })
			}
		}
		//}
	}

	return props
}



