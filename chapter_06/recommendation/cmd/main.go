package main

import (
	"DDD/chapter_06/recommendation/internal/recommendation"
	"DDD/chapter_06/recommendation/internal/transport"
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	c := retryablehttp.NewClient()
	c.RetryMax = 10

	partnerAdaptor, err := recommendation.NewPartnerShipAdaptor(c.StandardClient(), "http://localhost:3031",)
	if err != nil {
		log.Fatal("failed to create a partnerAdaptor")
	}

	svc, err := recommendation.NewService(partnerAdaptor)
	if err != nil {
		log.Fatal("failed to create a service: %w", err)
	}
	handler, err := recommendation.NewHandler(*svc)
	if err != nil {
		log.Fatal("failed to create a handler: %w", err)
	}

	m := transport.NewMux(*handler)
	if err := http.ListenAndServe(":4040", m); err != nil {
		log.Fatal("server errored: %w", err)
	}
}