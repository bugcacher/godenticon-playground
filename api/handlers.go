package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bugcacher/godenticon-playground/logger"
	"github.com/bugcacher/godenticon/avatar"
)

func HandleGenerateIdenticon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}
	queryParams := r.URL.Query()
	err := validateGenerateAvatar(queryParams)
	if err != nil {
		logger.DefaultLogger.Error("failed to validate query params", "error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var createOptions []avatar.CreateOption
	createOptions = append(createOptions, avatar.WithOutputType(avatar.OUTPUT_BUFFER))

	value := queryParams.Get(QueryParam_Value)

	if queryParams.Has(QueryParam_PixelPattern) {
		val, _ := strconv.Atoi(queryParams.Get(QueryParam_PixelPattern))
		createOptions = append(createOptions, avatar.WithPixelPattern(avatar.PixelPattern(val)))
	}
	if queryParams.Has(QueryParam_Algo) {
		val, _ := strconv.Atoi(queryParams.Get(QueryParam_Algo))
		createOptions = append(createOptions, avatar.WithAlgorithm(avatar.Algorithm(val)))
	}
	if queryParams.Has(QueryParam_DarkMode) {
		val, _ := strconv.ParseBool(queryParams.Get(QueryParam_DarkMode))
		if val {
			createOptions = append(createOptions, avatar.WithDarkMode())
		}
	}
	if queryParams.Has(QueryParam_Dimension) {
		val, _ := strconv.Atoi(queryParams.Get(QueryParam_Dimension))
		createOptions = append(createOptions, avatar.WithDimension(uint(val)))
	}

	avatar := avatar.New(value, createOptions...)
	res, err := avatar.Generate()
	if err != nil {
		logger.DefaultLogger.Error("failed to generate avatar", "error", err.Error())
		http.Error(w, ErrInterServerError.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res.Buffer.Bytes())
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(res.Buffer.Bytes())))
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
