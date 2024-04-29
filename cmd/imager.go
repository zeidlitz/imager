package main

import (
  "log/slog"
  "fmt"
  
  "github.com/zeidlitz/imager/internal/env"
	"github.com/zeidlitz/imager/internal/server"
)

type Config struct {
  serverAddress string
  httpPort int
}

func main() {
  var cfg Config
  cfg.serverAddress = env.GetString("SERVER_ADDRESS", "localhost")
  cfg.httpPort = env.GetInt("HTTP_PORT", 8080)
  address := fmt.Sprintf(cfg.serverAddress + ":" + "%d", 8080)
  slog.Info("Starting on", "address", address)
	server.Run(address)
}
