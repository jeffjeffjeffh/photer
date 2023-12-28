package main

import (
	"image"
	"log"
	"net/http"
	"sync"

	_ "image/jpeg"
)

type apiConfig struct{
	mut *sync.Mutex
	img image.Image
	newImg image.Image
}

func main() {
	const PORT string = "5001"
	mux := http.NewServeMux()
	apiCfg := apiConfig{
		mut: &sync.Mutex{},
	}

	mux.HandleFunc("/upload", apiCfg.handlerUpload)
	mux.HandleFunc("/resize", apiCfg.handlerResize)

	server := &http.Server{
		Handler: mux,
		Addr: ":" + PORT,
	}

	log.Printf("Server listening on port %s\n", PORT)
	server.ListenAndServe()
}