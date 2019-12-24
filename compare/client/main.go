package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func main() {
	t1 := time.Now()
	for i := 0; i < 1000; i++ {
		r, err := sayHello(fmt.Sprintf("yang %d", i))
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("hello mr {%s}", r))
		err = sayBye(fmt.Sprintf("yang %d", i))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(time.Since(t1))

}

func sayHello(name string) (string, error) {

	jsonStr, err := json.Marshal(HelloReq{Name: name})
	if err != nil {
		panic(err)
	}
	url := "http://localhost:8090/sayhello"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	r := HelloResp{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	return r.Message, nil
}
func sayBye(name string) error {

	jsonStr, err := json.Marshal(ByeReq{Name: name})
	if err != nil {
		panic(err)
	}
	url := "http://localhost:8090/saybye"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
