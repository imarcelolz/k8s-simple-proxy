package main

// import (
// 	"context"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	kube "k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/rest"
// )

// type ServiceCallback = func(service string)

// type Monitor struct {
// 	services         []string
// 	namespace        string
// 	kubeClient       *kube.Clientset
// 	onServiceAdded   ServiceCallback
// 	onServiceRemoved ServiceCallback
// }

// func CreateMonitor(namespace string, onServiceAdded ServiceCallback, onServiceRemoved ServiceCallback) *Monitor {
// 	// monitorContext := context.Background()

// 	config, _ := rest.InClusterConfig()
// 	kubeClient, _ := kube.NewForConfig(config)
// 	services, _ := findActivePods(kubeClient, namespace)

// 	instance := &Monitor{
// 		namespace:        namespace,
// 		kubeClient:       kubeClient,
// 		onServiceAdded:   onServiceAdded,
// 		onServiceRemoved: onServiceRemoved,
// 		services:         services,
// 	}

// 	//go watchKubeNamespace(instance, namespace)

// 	return instance
// }

// func findActivePods(client *kube.Clientset, namespace string) ([]string, error) {
// 	pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

// 	for i, pod := range pods.Items {

// 	}

// 	return []string{}, err
// }

// // func watchKubeNamespace(monitor *Monitor, namespace string) {
// // }
