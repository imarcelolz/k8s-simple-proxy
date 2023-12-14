package main

type Service = string
type ServiceChannel = chan Service

type ServiceMonitor interface {
	Services() []Service
	Watch(onServiceAdded ServiceChannel, onServiceRemoved ServiceChannel)
}
