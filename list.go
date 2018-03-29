package main

import (
	"fmt"
)

type ListCommand struct{}

func (l *ListCommand) Execute(args []string) error {
	triggers, err := getTriggers(options.APIHost, options.Dataset, options.WriteKey)
	if err != nil {
		fmt.Println("Failed to list triggers: ", err)
		return err
	}

	for _, trigger := range triggers {
		fmt.Println(trigger.Name)
	}

	return nil
}
