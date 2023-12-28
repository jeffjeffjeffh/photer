package main

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/disintegration/imaging"
)

func (apiCfg *apiConfig) handlerResize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("method not allowed")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	width, height, err := getParams(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(io.Discard, r.Body)
	if err != nil {
		log.Println("error discarding request body")
	}

	apiCfg.mut.Lock()
	defer apiCfg.mut.Unlock()

	if apiCfg.img == nil {
		err := errors.New("no image uploaded")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// params and image file are ok, do the resize
	resizedImg := imaging.Resize(apiCfg.img, width, height, imaging.NearestNeighbor)
	apiCfg.newImg = resizedImg
	log.Println("image resized")
	w.WriteHeader(http.StatusOK)
}