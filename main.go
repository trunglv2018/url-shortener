package main

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"url-shortener/config"
	"url-shortener/handlers"
	"url-shortener/storages"

	"github.com/mitchellh/go-homedir"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dir, _ := homedir.Dir()
	storage := &storages.Filesystem{}
	err := storage.Init(filepath.Join(dir, "shawty"))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handlers.EncodeHandler(storage))
	http.Handle("/dec/", handlers.DecodeHandler(storage))
	http.Handle("/red/", handlers.RedirectHandler(storage))

	port := config.GetConfig().GetString("server.port")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
