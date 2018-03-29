package main

type trigger struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Frequency   int                `json:"frequency,omitempty"`
	Query       *querySpec         `json:"query,omitempty"`
	Threshold   *triggerThreshold  `json:"threshold,omitempty"`
	Recipients  []triggerRecipient `json:"recipients,omitempty"`

	ID string `json:"id,omitempty"`
}

type triggerRecipient struct {
	Type   string `json:"type"`
	Target string `json:"target,omitempty"`
}

type triggerThreshold struct {
	Op    string   `json:"op"`
	Value *float64 `json:"value"`
}

type querySpec struct {
	Breakdowns        []string        `json:"breakdowns,omitempty"`
	Calculations      []calculateSpec `json:"calculations,omitempty"`
	Filters           []filterSpec    `json:"filters,omitempty"`
	FilterCombination *string         `json:"filter_combination,omitempty"`
}

type calculateSpec struct {
	Column *string `json:"column,omitempty"`
	Op     string  `json:"op"`
}

type filterSpec struct {
	Column string      `json:"column"`
	Op     string      `json:"op"`
	Value  interface{} `json:"value,omitempty"`
}

type configFile struct {
	Triggers []trigger `json:"triggers"`
}
