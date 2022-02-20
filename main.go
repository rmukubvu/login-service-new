package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"login-service/config"
	"login-service/controller"
	"login-service/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port = flag.Int("p", config.WebServerPort(), "port number")

func main() {

	flag.Parse()
	r := controller.InitRouter()
	sPort := fmt.Sprintf(":%d", *port)

	srv := &http.Server{
		Addr: sPort,
		Handler: handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r),
	}
	//show on stdout
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started. Press Ctrl-C to stop server")
	// Catch the Ctrl-C and SIGTERM from kill command
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//done
	log.Println("Shutting down server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	service.TearDown()
}
