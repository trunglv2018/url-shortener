package main

import (
	//1.load firstly

	"url-shortener/config"
	"url-shortener/middleware"
	"url-shortener/model"

	//2.
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"url-shortener/api"

	"github.com/gin-gonic/gin"
	"github.com/trunglen/g/x/rest"
)

// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())

// 	dir, _ := homedir.Dir()
// 	storage := &storages.Filesystem{}
// 	err := storage.Init(filepath.Join(dir, "shawty"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	http.Handle("/", handlers.EncodeHandler(storage))
// 	http.Handle("/dec/", handlers.DecodeHandler(storage))
// 	http.Handle("/red/", handlers.RedirectHandler(storage))

// 	port := config.GetConfig().GetString("server.port")
// 	err = http.ListenAndServe(":"+port, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
func main() {
	router := gin.Default()
	router.Use(middleware.Recovery())
	api.NewApi(router.Group("api/v1"))
	router.GET("/p/:id", payloadHandler)
	router.GET("/:id", redirectHandler)
	srv := &http.Server{
		Addr:    ":" + config.GetConfig().GetString("server.port"),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}

func redirectHandler(c *gin.Context) {
	var linkID, _ = c.Params.Get("id")
	var shortLink, err = new(model.Link).GetByCode(linkID)
	rest.AssertNil(err)
	rest.AssertNil(shortLink.Visit())
	c.Redirect(301, shortLink.LongLink)
}

func payloadHandler(c *gin.Context) {
	var linkID, _ = c.Params.Get("id")
	var shortLink, err = new(model.Link).GetByCode(linkID)
	rest.AssertNil(err)
	c.JSON(200, map[string]interface{}{
		"data":   shortLink.Payload,
		"status": "success",
	})
}
