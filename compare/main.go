package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloReq struct {
	Name string `json:"name"`
}
type HelloResp struct {
	Message string `json:"message"`
}
type ByeReq struct {
	Name string `json:"name"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	p := HelloReq{}
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := HelloResp{Message: fmt.Sprintf("hello %s", p.Name)}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}
}

func bye(w http.ResponseWriter, req *http.Request) {

	p := ByeReq{}
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(p.Name)
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/sayhello", hello)
	http.HandleFunc("/saybye", bye)
	http.ListenAndServe(":8090", nil)
}
