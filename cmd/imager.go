package main

import (
  "encoding/json"
  "net/http"
)

var(
  count int
)

func getJsonResponseBytes(status int, detail string) ([]byte, error) {
  response := map[string]string{"status": status,  "detail" : detail}
	jsonResponse, _ := json.Marshal(response)
  return jsonResponse, nil
}

func indexPage(w http.ResponseWriter, req *http.Request) {
  if(req.Method != http.MethodGet) {
    statusCode := http.StatusMethodNotAllowed
    detail := "Invalid method"
    msg, _ := getJsonResponseBytes(statusCode, "Invalid method")
    http.Error(w, detail, http.StatusMethodNotAllowed) 
    w.Write(msg)
    return
  }
  count++
}

func main() {
  http.HandleFunc("/", indexPage)
}
