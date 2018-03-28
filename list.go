package main

import (
	"encoding/json"
	"fmt"
)

type ListCommand struct{}

func (l *ListCommand) Execute(args []string) error {
	triggers, err := getTriggers(options.APIHost, options.Dataset, options.WriteKey)
	if err != nil {
		fmt.Println("Failed to list triggers: ", err)
		return err
	}

	data, err := json.Marshal(triggers)
	if err != nil {
		fmt.Println("Failed to list triggers: ", err)
		return err
	}
	fmt.Println(string(data))

	return nil
}
