package main

import (
	"github.com/gorilla/mux"
	"github.com/vitkarpov/lorempixum"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlerIndex)
	router.HandleFunc("/{format}/{width}/{height}", handlerImage)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handlerIndex(res http.ResponseWriter, req *http.Request) {
	http.ServeContent(res, req, "static/index.html")
}

func handlerImage(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	width, err := strconv.Atoi(params["width"])
	height, err := strconv.Atoi(params["height"])
	size := width * height

	if err != nil || size <= 0 || size > 1000000 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	img := lorempixum.GetImage(width, height)
	lorempixum.StreamImage(res, img, params["format"])
}
