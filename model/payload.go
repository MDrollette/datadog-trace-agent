package model

// AgentPayload is the payload sent to the API of mothership with raw traces
type AgentPayload struct {
	HostName string      `json:"hostname"`
	Spans    []Span      `json:"spans"`
	Stats    StatsBucket `json:"stats"`
}

func (p *AgentPayload) IsEmpty() bool {
	return len(p.Spans) == 0 && p.Stats.IsEmpty()
}
