package kubedata

import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getConfig() (clientset *kubernetes.Clientset) {
	dirname, _ := os.UserHomeDir()

	config, _ := clientcmd.BuildConfigFromFlags("", dirname+"/.kube/config")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return clientset
}
