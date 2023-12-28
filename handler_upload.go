package main

import (
	"errors"
	"image"
	"log"
	"net/http"
)

func (apiCfg *apiConfig) handlerUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("method not allowed")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	apiCfg.mut.Lock()
	defer apiCfg.mut.Unlock()

	parseMem := int64(10 << 20)
	err := r.ParseMultipartForm(parseMem)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("file extracted from form")
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("image decoded")

	apiCfg.img = img
	log.Println(apiCfg.img.Bounds())
}