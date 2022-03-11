package main

import (
	"Technopark_3_Security/proxyHandlers"
	"Technopark_3_Security/repositories"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"

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

func (repeater *Repeater) HandleScanRequest(w http.ResponseWriter, req *http.Request) {
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

	if res.Method != http.MethodGet {
		http.Error(w, "Bad method type", http.StatusBadRequest)
		return
	}

	resReq, err := res.GetHTTPRequest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	inputSource, err := os.Open(paramsFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer func(inputSource *os.File) {
		_ = inputSource.Close()
	}(inputSource)

	scanner := bufio.NewScanner(inputSource)
	var params []string
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}

	someRandValue := uuid.NewString()
	for _, param := range params {
		resReqWithParam := resReq
		query := resReqWithParam.URL.Query()
		query.Add(param, someRandValue)
		resReqWithParam.URL.RawQuery = query.Encode()

		response, err := http.DefaultTransport.RoundTrip(resReqWithParam)
		if err != nil {
			fmt.Println(err)
			return
		}

		if response.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(response.Body)
			_ = response.Body.Close()
			if err != nil {
				continue
			}
			bodyString := string(bodyBytes)
			if strings.Contains(bodyString, param) {
				fmt.Println(param)
			}
		} else {
			_ = response.Body.Close()
		}
	}

	proxyHandlers.HandleHTTP(w, resReq)
}
