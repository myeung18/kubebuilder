package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func main() {
	start()
}

func start() {
	kubeconfig := flag.String("kubeconfig", "/home/myeung/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		os.Exit(1)
	}
	clientset, _ := kubernetes.NewForConfig(config)

	pod, err := clientset.CoreV1().Pods("mss-test").Get(context.TODO(), "example", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("The clientset cannot be loaded: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(pod.Spec)
}
