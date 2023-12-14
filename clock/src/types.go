package main

import (
	"encoding/json"
	"fmt"
)

type ClockState struct {
	id        string
	runing    bool
	precision uint16
	timeLeft  uint16
}

func encodeClock(state *ClockState) (string, error) {
	buffer, err := json.Marshal(state)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(buffer), nil
}

func decodeClock(serialized string) (*ClockState, error) {
	buffer := new(ClockState)
	err := json.Unmarshal([]byte(serialized), buffer)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return buffer, nil
}
