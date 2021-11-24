package main

import (
	"fmt"

	"github.com/michalwr/k8status/internal/kubedata"
)
func main() {

	Kube:=kubedata.InitKubernetes()
	var props []kubedata.PodProps

    props=append(props, kubedata.PodProps{Name:"test"})
	fmt.Println( Kube.PodProps)
	Kube.PodProps=props
	Kube.RefreshData()
	fmt.Println( Kube.PodProps)

}