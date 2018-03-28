package main

type trigger struct {
	ID string `json:"id,omitempty"`

	Threshold threshold `json:"threshold"`

	Description string      `json:"description,omitempty"`
	Frequency   int64       `json:"frequency"`
	Name        string      `json:"name"`
	Recipients  []recipient `json:"recipients"`
	Query       query       `json:"query"`
}

type recipient struct {
	Type   string `json:"type"`
	Target string `json:"target"`
}

type threshold struct {
	Value int    `json:"value"`
	Op    string `json:"op"`
}

type query struct {
	Calculations []calculation `json:"calculations"`
	Filters      []filter      `json:"filters"`
	Breakdowns   []string      `json:"breakdowns"`
}

type calculation struct {
	Op string `json:"op"`
}

type filter struct {
	Value  string `json:"value"`
	Op     string `json:"op"`
	Column string `json:"column"`
}
