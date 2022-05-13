package models

import (
	"TheWarEconomy/api/utils"

	"github.com/google/uuid"
)

type Skills struct {
	Combat        int64 `json:"combat"`
	Research      int64 `json:"research"`
	Medical       int64 `json:"medical"`
	Intel         int64 `json:"intel"`
	Food          int64 `json:"food"`
	Transport     int64 `json:"transport"`
	Maintenance   int64 `json:"maintenance"`
	Manufacturing int64 `json:"manufacturing"`
}

type Soldier struct {
	Id     string   `json:"id"`
	Skills Skills   `json:"skills"`
	Traits []string `json:"traits"`
	Name   string   `json:"name"`
}

const GENERATED_SOLDIER_MIN_SKILL = 1
const GENERATED_SOLDIER_MAX_SKILL = 80

func GenerateSoldiers(amount int) []Soldier {
	var soldiers []Soldier

	for i := 0; i < amount; i++ {
		var soldier Soldier
		soldier.Id = uuid.NewString()
		soldier.Skills.Combat = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Skills.Research = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Skills.Medical = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Skills.Intel = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Skills.Food = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Skills.Transport = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Skills.Maintenance = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Skills.Manufacturing = utils.GenerateNumberInRange(GENERATED_SOLDIER_MIN_SKILL, GENERATED_SOLDIER_MAX_SKILL)
		soldier.Traits = make([]string, 0)
		soldier.Name = "Terrifying Potato"
		soldiers = append(soldiers, soldier)
	}

	return soldiers
}
