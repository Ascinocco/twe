package controllers

import (
	"TheWarEconomy/api/models"
	"TheWarEconomy/api/utils"
	"encoding/json"
	"net/http"
)

type PmcResponse struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Error  string `json:"error"`
}

func CreatePmc(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIdFromCtx(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&PmcResponse{
			Error: "Unable to create PMC, please try again",
		})
		return
	}

	pmc := &models.Pmc{}
	err = json.NewDecoder(r.Body).Decode(pmc)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&PmcResponse{
			Error: "Unable to create PMC, please try again",
		})
		return
	}

	p, err := pmc.Create(userId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&PmcResponse{
			Error: "Unable to create PMC, please try again",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&PmcResponse{
		Id:     p.Id,
		UserId: p.UserId,
		Name:   p.Name,
	})
}
