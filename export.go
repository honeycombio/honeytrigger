package main

import (
	"encoding/json"
	"fmt"
)

type ExportCommand struct{}

func (e *ExportCommand) Execute(args []string) error {
	triggers, err := getTriggers(options.APIHost, options.Dataset, options.WriteKey)
	if err != nil {
		fmt.Println("Failed to get triggers: ", err)
		return err
	}

	for i := range triggers {
		// strip the ID when exporting a config
		triggers[i].ID = ""
	}

	config, err := json.Marshal(configFile{Triggers: triggers})
	if err != nil {
		fmt.Println("Failed to build config: ", err)
		return err
	}
	fmt.Println(string(config))

	return nil
}
