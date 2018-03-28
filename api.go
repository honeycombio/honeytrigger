package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getTriggers(apiHost, dataset, writeKey string) ([]trigger, error) {
	postURL, err := url.Parse(apiHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s", options.APIHost)
		return nil, errors.New(errMsg)
	}
	postURL.Path = "/1/triggers/" + dataset
	req, err := http.NewRequest("GET", postURL.String(), nil)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Add("X-Honeycomb-Team", writeKey)
	req.Header.Add("X-Honeycomb-Dataset", dataset)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Failed with %d and message: %s", resp.StatusCode, body)
		return nil, errors.New(errMsg)
	}

	var triggers []trigger
	err = json.Unmarshal(body, &triggers)
	if err != nil {
		return nil, err
	}

	return triggers, nil
}

func addTrigger(apiHost, dataset, writeKey string, t *trigger) error {
	postURL, err := url.Parse(apiHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s", options.APIHost)
		return errors.New(errMsg)
	}
	blob, err := json.Marshal(*t)
	if err != nil {
		return err
	}
	postURL.Path = "/1/triggers/" + dataset
	req, err := http.NewRequest("POST", postURL.String(), bytes.NewBuffer(blob))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Add("X-Honeycomb-Team", writeKey)
	req.Header.Add("X-Honeycomb-Dataset", dataset)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		errMsg := fmt.Sprintf("Failed with %d and message: %s", resp.StatusCode, body)
		return errors.New(errMsg)
	}

	return nil
}

func updateTrigger(apiHost, dataset, writeKey string, t *trigger) error {
	postURL, err := url.Parse(apiHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s", options.APIHost)
		return errors.New(errMsg)
	}
	blob, err := json.Marshal(*t)
	if err != nil {
		return err
	}
	postURL.Path = "/1/triggers/" + dataset + "/" + t.ID
	req, err := http.NewRequest("PUT", postURL.String(), bytes.NewBuffer(blob))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Add("X-Honeycomb-Team", writeKey)
	req.Header.Add("X-Honeycomb-Dataset", dataset)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Failed with %d and message: %s", resp.StatusCode, body)
		return errors.New(errMsg)
	}

	return nil
}
