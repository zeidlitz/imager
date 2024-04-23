package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getJsonResponseBytes(detail string, code int) ([]byte, error) {
  statusCode := fmt.Sprintf("%d", code)
  response := map[string]string{"status": statusCode,  "detail" : detail}
	jsonResponse, _ := json.Marshal(response)
  return jsonResponse, nil
}

func indexPage(w http.ResponseWriter, req *http.Request) {
  if(req.Method != http.MethodGet) {

    detail := "Invalid method"
    statusCode := http.StatusMethodNotAllowed

    http.Error(w, detail, http.StatusMethodNotAllowed) 
    msg, _ := getJsonResponseBytes(detail, statusCode)
    w.Write(msg)

    return
  }

  detail := "OK"
  statusCode := http.StatusOK
  msg, _ := getJsonResponseBytes(detail, statusCode)

  w.Write(msg)
}

func main() {
  address := "localhost:8080"
  fmt.Println("Listening on ", address)
  http.HandleFunc("/", indexPage)
  http.ListenAndServe(address, nil)
}
