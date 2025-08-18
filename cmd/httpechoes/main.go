package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gulien/http-echoes/pkg/httpechoes"
)

func main() {
	config := httpechoes.NewDefaultConfig()
	server := httpechoes.New(config)

	fmt.Printf("[SYSTEM] Listening on %s\n", config.Addr())
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("[FATAL] %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
