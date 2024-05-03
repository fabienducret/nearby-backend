package informations

import (
	"encoding/json"
	"nearby/models"
	"net/http"
)

type GetInformationsUseCase func(city string) models.Informations

func InitController(getInfos GetInformationsUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		city := r.PathValue("city")
		infos := getInfos(city)

		response, _ := json.Marshal(infos)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
