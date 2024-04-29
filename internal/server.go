package server

import (
	"encoding/json"
	"fmt"
  "os"
	"net/http"
)

type PageData struct {
  Paragraph string
}

func getJsonResponseBytes(detail string, code int) ([]byte, error) {
  statusCode := fmt.Sprintf("%d", code)
  response := map[string]string{"status": statusCode,  "detail" : detail}
	jsonResponse, _ := json.Marshal(response)
  return jsonResponse, nil
}

func invalidMethod(w http.ResponseWriter){
    detail := "Invalid method"
    statusCode := http.StatusMethodNotAllowed

    http.Error(w, detail, http.StatusMethodNotAllowed) 
    msg, _ := getJsonResponseBytes(detail, statusCode)
    w.Write(msg)

    return
}

func faviconHandler(w http.ResponseWriter, req *http.Request) {
  if(req.Method != http.MethodGet) {
    invalidMethod(w)
  }

  favicon, err := os.ReadFile("assets/favicon.png")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  w.Write(favicon)
}

func logoHandler(w http.ResponseWriter, req *http.Request) {
  if(req.Method != http.MethodGet) {
    invalidMethod(w)
  }

  logo, err := os.ReadFile("assets/logo.png")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  w.Write(logo)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
  if(req.Method != http.MethodGet) {
    invalidMethod(w)
  }

  html, err := os.ReadFile("web/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  w.Write(html)

}

func run() {
  address := "localhost:8080"
  fmt.Println("Listening on ", address)
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/logo", logoHandler)
  http.HandleFunc("/favicon", faviconHandler)
  http.ListenAndServe(address, nil)
}
