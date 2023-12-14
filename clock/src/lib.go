package main

import "sync"

type Action func()

func lock(mutex *sync.RWMutex, action Action) {
	mutex.Lock()
	action()
	mutex.Unlock()
}
