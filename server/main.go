/*
 * openapi server
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	OkaserverAPIService := openapi.NewOkaserverAPIService()
	OkaserverAPIController := openapi.NewOkaserverAPIController(OkaserverAPIService)

	router := openapi.NewRouter(OkaserverAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
