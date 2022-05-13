package models

import (
	"TheWarEconomy/api/database"
	"errors"

	"github.com/couchbase/gocb/v2"
)

type Pmc struct {
	Id          string      `json:"id"`
	UserId      string      `json:"userId"`
	Name        string      `json:"name"`
	Soliders    []Soldier   `json:"soldiers"`
	Money       int64       `json:"money"`
	Deployments Deployments `json:"deployments"`
}

func (pmc *Pmc) Validate() (string, bool) {
	if len(pmc.Name) <= 5 {
		return "PMC name must be at least 5 characters", false
	}

	return "", true
}

func (pmc *Pmc) Create(userId string) (*Pmc, error) {
	if errMsg, ok := pmc.Validate(); !ok {
		return pmc, errors.New(errMsg)
	}

	pmc.Id = pmc.Name
	pmc.UserId = userId
	pmc.Money = 10000000 // give 100k on creation
	pmc.Soliders = GenerateSoldiers(50)
	pmcc := database.GetPmcsCol()
	_, err := pmcc.Insert(pmc.Id, pmc, nil)

	if err == nil {
		uc := database.GetUsersCol()
		upds := []gocb.MutateInSpec{
			gocb.UpsertSpec("pmcId", pmc.Id, nil),
		}
		uc.MutateIn(userId, upds, nil)
	}

	return pmc, err
}
