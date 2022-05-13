package models

type Rewards struct {
	Money int64 `json:"money"`
}

type Missions struct {
	Type          string    `json:"type"`
	EnemySoldiers []Soldier `json:"enemySoldiers"`
	Rewards       Rewards   `json:"rewards"`
	Duration      int64     `json:"duration"`
}
