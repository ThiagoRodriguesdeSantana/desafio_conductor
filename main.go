/*
 * Desafio conductor
 *
 * Api para controle de trasacoes de contas
 *
 * API version: 1.0.0
 * Contact: thiagorodriguescamara@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "github.com/ThiagoRodriguesdeSantana/desafio_conductor/go"
)

func main() {

	pathDb := os.Args[1]

	fmt.Println(pathDb)
	log.Printf("Server started")

	router := sw.NewRouter(pathDb)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8181",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Print("127.0.0.1:8181 Running...")
	log.Fatal(srv.ListenAndServe())

}