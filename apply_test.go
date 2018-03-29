package main

import (
	"testing"
)

func TestDiffTriggers(t *testing.T) {
	existingTriggers := []trigger{
		trigger{Name: "foo"},
	}
	desiredTriggers := []trigger{
		trigger{Name: "foo"},
		trigger{Name: "bar"},
	}

	newTriggers, currentTriggers := diffTriggers(existingTriggers, desiredTriggers)
	if len(newTriggers) != 1 {
		t.Errorf("unexpected number of newTriggers returned. Expected 1, got %d", len(newTriggers))
	}
	if len(currentTriggers) != 1 {
		t.Errorf("unexpected number of currentTriggers returned. Expected 1, got %d", len(currentTriggers))
	}
	if newTriggers[0].Name != "bar" {
		t.Errorf("unexpected value in newTriggers. Expected bar, got %s", newTriggers[0].Name)
	}
	if currentTriggers[0].Name != "foo" {
		t.Errorf("unexpected value in currentTriggers. Expected bar, got %s", currentTriggers[0].Name)
	}
}
