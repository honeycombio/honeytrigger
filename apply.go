package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configFile struct {
	Triggers []trigger `json:"triggers"`
}

type ApplyCommand struct {
	File string `short:"f" long:"config_file" description:"Path to config file for the given dataset" required:"true"`
}

func (a *ApplyCommand) Execute(args []string) error {
	data, err := ioutil.ReadFile(a.File)
	if err != nil {
		fmt.Println("Failed to read config file: ", a.File)
		return err
	}
	var config configFile
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Failed to parse config file: ", err)
	}

	existingTriggers, err := getTriggers(options.APIHost, options.Dataset, options.WriteKey)
	if err != nil {
		fmt.Println("Failed to get existing triggers: ", err)
		return err
	}

	newTriggers, currentTriggers := diffTriggers(existingTriggers, config.Triggers)

	for i := range newTriggers {
		fmt.Printf("Adding trigger '%s'\n", newTriggers[i].Name)
		err = addTrigger(options.APIHost, options.Dataset, options.WriteKey, &newTriggers[i])
		if err != nil {
			fmt.Printf("Failed to add new trigger '%s': %s", newTriggers[i].Name, err)
			return err
		}
	}

	for i := range currentTriggers {
		fmt.Printf("Updating trigger '%s' with id %s\n", currentTriggers[i].Name, currentTriggers[i].ID)
		err = updateTrigger(options.APIHost, options.Dataset, options.WriteKey, &currentTriggers[i])
		if err != nil {
			fmt.Printf("Failed to update trigger '%s': %s", newTriggers[i].Name, err)
			return err
		}
	}

	return nil
}

// diffTriggers compares current triggers in the API with what is in the config
func diffTriggers(existingTriggers, desiredTriggers []trigger) ([]trigger, []trigger) {
	var newTriggers, currentTriggers []trigger
	existingNames := make(map[string]string)

	for i := range existingTriggers {
		existingNames[existingTriggers[i].Name] = existingTriggers[i].ID
	}

	for i := range desiredTriggers {
		if id, ok := existingNames[desiredTriggers[i].Name]; ok {
			// map the ID for the existing trigger to the name, so we can use it
			// when we do a PUT later
			desiredTriggers[i].ID = id
			currentTriggers = append(currentTriggers, desiredTriggers[i])
		} else {
			newTriggers = append(newTriggers, desiredTriggers[i])
		}
	}

	return newTriggers, currentTriggers
}
