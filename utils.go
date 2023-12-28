package main

import (
	"errors"
	"net/http"
	"strconv"
)

func getParams(r *http.Request) (int, int, error) {
	width := r.URL.Query().Get("width")
	height := r.URL.Query().Get("height")

	if width == "" && height == "" {
		return 0, 0, errors.New("no width or height specified")
	}

	var widthInt, heightInt int

	if width == "" {
		widthInt = 0
	} else {
		parsedInt, err := strconv.Atoi(width)
		if err != nil {
			return 0, 0, errors.New("bad width param")
		}
		widthInt = parsedInt
	}

	if height == "" {
		heightInt = 0
	} else {
		parsedInt, err := strconv.Atoi(height)
		if err != nil {
			return 0, 0, errors.New("bad height param")
		}
		heightInt = parsedInt
	}

	return widthInt, heightInt, nil
}