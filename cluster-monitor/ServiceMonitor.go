package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/1.5/rest"
	kube "k8s.io/client-go/kubernetes"
)

type Service = string
type ServiceChannel = chan Service

type ServiceMonitor interface {
	Services() []Service
	Watch(onServiceAdded ServiceChannel, onServiceRemoved ServiceChannel)
}
type Monitor struct {
	services   []Service
	namespace  string
	kubeClient *kube.Clientset
}

func (monitor *Monitor) Watch(onServiceAdded ServiceChannel, onServiceRemoved ServiceChannel) {}
func (monitor *Monitor) Services() []Service {
	return monitor.services
}

func NewServiceMonitor(namespace string) ServiceMonitor {
	config, _ := rest.InClusterConfig()
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
	pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic("AA")
	}

	podNames := make([]string, pods.Size())
	for i, pod := range pods.Items {
		podNames[i] = pod.Name
	}

	return podNames
}
