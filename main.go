package main

import (
	"errors"
	"image"
	"log"
	"net/http"
)

func main() {
	const PORT string = "5001"
	mux := http.NewServeMux()

	mux.HandleFunc("/resize", handleResize)

	server := &http.Server{
		Handler: mux,
		Addr: ":" + PORT,
	}

	log.Printf("Server listening on port %s\n", PORT)
	server.ListenAndServe()
}

func handleResize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("method not allowed")
		log.Println(err)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	log.Println("request content-type:", r.Header.Get("Content-Type"))

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Fatal("Error getting image from form:", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Failed to decode the image:", err)
	}

	log.Println(img.Bounds())

	w.WriteHeader(http.StatusOK)
}