package app

import (
	"fmt"
	"go-flashcard/app/config"
	"go-flashcard/app/router"
	"log"
	"net/http"
)

func Run() {
	fmt.Printf("\n\tListening[::]:%d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
