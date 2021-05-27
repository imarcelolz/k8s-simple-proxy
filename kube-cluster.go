package main

import (
	"context"

	metav1 "g/pkg/apis/meta/v1"

	kubeRest "k8s.io/client-go/1.5/rest"
	kube "k8s.io/client-go/kubernetes"
)

type ServiceCallback = func(service string)

type Monitor struct {
	services         []string
	namespace        string
	kubeClient       *kube.Clientset
	onServiceAdded   ServiceCallback
	onServiceRemoved ServiceCallback
}

func CreateMonitor(namespace string, onServiceAdded ServiceCallback, onServiceRemoved ServiceCallback) *Monitor {
	// monitorContext := context.Background()

	kubeClient, _ := kube.NewForConfig(kubeRest.InClusterConfig())
	services, _ := findActivePods(kubeClient, namespace)

	instance := &Monitor{
		namespace:        namespace,
		kubeClient:       kubeClient,
		onServiceAdded:   onServiceAdded,
		onServiceRemoved: onServiceRemoved,
		services:         services,
	}

	go watchKubeNamespace(instance, namespace)

	return instance
}

func findActivePods(client *kube.Clientset, namespace string) ([]string, error) {
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	podList.Items
	return []string{}, err
}

func watchKubeNamespace(monitor *Monitor, namespace string) {
}
