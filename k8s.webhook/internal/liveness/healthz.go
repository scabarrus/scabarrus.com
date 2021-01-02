package liveness

import (
	"encoding/json"
	"net/http"

	format "scabarrus.com/k8s.webhook/internal/error"

	"scabarrus.com/k8s.webhook/internal/repository"
)

// Healthz godoc
// @Summary healthcheck
// @Description send OK if it database connexion works
// @Tags healthz
// @Accept  json
// @Produce  json
// @Success 200 {object} string true "status"
// @Success 400 {object} format.Error
// @Router /healthz [get]
func Healthz(w http.ResponseWriter,r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	err:= pg.Healthz()
	var e format.Error
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		e.FormatError("SQL Error - ",err.Error(),r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		
	}else{
		w.WriteHeader(http.StatusOK)
    	w.Write([]byte("ok"))

	}


	
}