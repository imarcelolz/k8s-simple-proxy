package main

type ClockInterface interface {
	create(clock *ClockState, precision uint16)
	find(id string) *ClockState

	start(id string) *ClockState
	stop(id string) *ClockState
	edit(id string, value uint16) *ClockState
}

type ClockStore struct {
	storage ClockStorage
}

func (this *ClockStore) create(initialState *ClockState) {
	id := initialState.id
	initialState.runing = false

	this.storage.set(id, initialState)
}

func (this *ClockStore) get(id string) *ClockState {
	return this.storage.get(id)
}

func (this *ClockStore) start(id string) *ClockState {
	mutation := func(state *ClockState) *ClockState {
		state.runing = true
		return state
	}

	return this.storage.update(id, mutation)
}

func (this *ClockStore) stop(id string) *ClockState {
	mutation := func(state *ClockState) *ClockState {
		state.runing = false
		return state
	}

	return this.storage.update(id, mutation)
}

func (this *ClockStore) edit(id string, timeLeft uint16) *ClockState {
	mutation := func(state *ClockState) *ClockState {
		state.timeLeft = timeLeft
		return state
	}

	return this.storage.update(id, mutation)
}

func (this *ClockStore) increment(id string, increment uint16) *ClockState {
	mutation := func(state *ClockState) *ClockState {
		state.timeLeft += increment
		return state
	}

	return this.storage.update(id, mutation)
}
