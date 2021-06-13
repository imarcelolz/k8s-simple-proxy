package main

import "sync"

type Mutation = func(state *ClockState) *ClockState

type ClockStorage interface {
	create() *ClockStorage

	get(sessionId string) *ClockState
	set(sessionId string, clock *ClockState)
	update(id string, mutation Mutation) *ClockState
	delete(sessionId string)
}

type ClockMemoryStorage struct {
	mutex   sync.RWMutex
	storage map[string]*ClockState
	locks   map[string]sync.Mutex
}

func createClockMemoryStorage() *ClockMemoryStorage {
	storage := new(ClockMemoryStorage)

	return storage
}

func (this *ClockMemoryStorage) update(id string, mutation Mutation) *ClockState {
	state, lock := __find(this, id)

	lock.Lock()
	updatedState := mutation(state)
	this.set(id, updatedState)
	lock.Unlock()

	return updatedState
}

func (this *ClockMemoryStorage) get(id string) *ClockState {
	this.mutex.RLock()
	clock := this.storage[id]
	this.mutex.RUnlock()

	return clock
}

func (this *ClockMemoryStorage) set(id string, clock *ClockState) {
	lock(&this.mutex, func() {
		clock.id = id
		this.storage[id] = clock
	})
}

func (this *ClockMemoryStorage) delete(id string) {
	lock(&this.mutex, func() {
		delete(this.storage, id)
		delete(this.locks, id)
	})
}

func __find(storage *ClockMemoryStorage, id string) (*ClockState, *sync.Mutex) {
	clockState := storage.storage[id]
	lock := storage.locks[id]

	return clockState, &lock
}
