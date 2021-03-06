package main

import (
	"github.com/gorilla/mux"
	"github.com/vitkarpov/lorempixum"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{format}/{width}/{height}", handlerImage)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
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
