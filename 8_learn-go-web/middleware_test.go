package learn_go_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(writer, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Executed")
		panic("Ups")
	})
	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	errorhandler := new(ErrorHandler)
	errorhandler.Handler = logMiddleware

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorhandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
