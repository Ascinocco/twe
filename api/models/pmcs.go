package models

import (
	"TheWarEconomy/api/database"
	"errors"
)

type Pmc struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func (pmc *Pmc) Validate() (string, bool) {
	if len(pmc.Name) <= 5 {
		return "PMC name must be at least 5 characters", false
	}

	return "", true
}

func (pmc *Pmc) Create() (*Pmc, error) {
	if errMsg, ok := pmc.Validate(); !ok {
		return pmc, errors.New(errMsg)
	}

	pmc.Id = pmc.Name
	pmcc := database.GetPmcsCol()
	_, err := pmcc.Insert(pmc.Id, pmc, nil)

	return pmc, err
}

func (pmc *Pmc) LinkUser(userId string) (*Pmc, error) {
	pmc.UserId = userId
	pmcc := database.GetPmcsCol()
	_, err := pmcc.Upsert(pmc.Id, pmc, nil)
	return pmc, err
}
