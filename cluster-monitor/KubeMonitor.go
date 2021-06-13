package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Monitor struct {
	services   []Service
	namespace  string
	kubeClient *kube.Clientset
}

func (monitor *Monitor) Watch(onServiceAdded ServiceChannel, onServiceRemoved ServiceChannel) {
	monitor.kubeClient.CoreV1().Pods().Watch(context.Background(), metav1.ListOptions{LabelSelector: "name"})

}
func (monitor *Monitor) Services() []Service {
	return monitor.services
}

func NewKubeMonitor(namespace string) ServiceMonitor {
	// config, err := rest.InClusterConfig()
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/marcelo/.kube/config.local")
	if err != nil {
		panic(err.Error())
	}

	kubeClient := kube.NewForConfigOrDie(config)

	services := findActivePods(kubeClient, namespace)

	serviceMonitor := &Monitor{
		namespace:  namespace,
		kubeClient: kubeClient,
		services:   services,
	}

	return serviceMonitor
}

func findActivePods(client *kube.Clientset, namespace string) []string {
	pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: "name"})

	if err != nil {
		panic(err)
	}

	podNames := make([]string, pods.Size())
	for i, pod := range pods.Items {
		podNames[i] = pod.Name
	}

	return podNames
}
