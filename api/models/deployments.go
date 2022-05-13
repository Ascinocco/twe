package models

type Deployments struct {
	Id          string    `json:"id"`
	Soldiers    []Soldier `json:"soldiers"`
	StartedAt   int64     `json:"startedAt"`
	CompletedAt int64     `json:"completedAt"`
	Type        string    `json:"type"` // e.g combat, intel, resource, procurment, etc
}
