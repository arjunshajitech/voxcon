package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"voxcon/game"
	"voxcon/space"
	"voxcon/util"
)

type JWTToken struct {
	Token string `json:"token"`
}

func Start(space *space.Space) {
	SetupRoutes()
	fmt.Println("Server started on port 7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		return
	}
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"health": "ok"}`))
	if err != nil {
		return
	}
}

func NewGame(w http.ResponseWriter, _ *http.Request) {
	id := util.GenerateRandomID()
	game.NewGame(id)
	token, err := util.CreateToken(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(&JWTToken{Token: token})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
