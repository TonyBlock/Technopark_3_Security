package main

import (
	"Technopark_3_Security/proxyHandlers"
	"Technopark_3_Security/repositories"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Repeater struct {
	Store *repositories.RequestStore
}

func (repeater *Repeater) HandleGetRequests(w http.ResponseWriter, req *http.Request) {
	res, err := repeater.Store.Get(20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

func (repeater *Repeater) HandleRepeatRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("HandleRepeatRequest ", req.RequestURI)
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := repeater.Store.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	resReq, err := res.GetHTTPRequest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyHandlers.HandleHTTP(w, resReq)
}

type Info struct {
	Status  int    `json:"status"`
	Path    string `json:"path"`
	Content string `json:"content"`
}
