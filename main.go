package main

import (
	integrase "github.com/MetroReviews/metro-integrase/lib"
	"github.com/MetroReviews/metro-integrase/types"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"
	"net/http"
)

// Dummy adapter backend
type DummyAdapter struct {
}

func (adp DummyAdapter) GetConfig() types.ListConfig {
	return types.ListConfig{
		SecretKey:   "GE6DWnDfKEgDrjNcdKvOIDgmtJ1KNtzoAjXfNrKF5wU",
		ListID:      "008900c1-96c0-4fba-82f9-e4f0ba904d73",
		RequestLogs: true,
		StartupLogs: true,
		DomainName:  "https://spider.infinitybotlist.com",
	}
}

func (adp DummyAdapter) ClaimBot(bot *types.Bot) error {
	return nil
}

func (adp DummyAdapter) UnclaimBot(bot *types.Bot) error {
	return nil
}

func (adp DummyAdapter) ApproveBot(bot *types.Bot) error {
	return nil
}

func (adp DummyAdapter) DenyBot(bot *types.Bot) error {
	return nil
}

func (adp DummyAdapter) DataDelete(id string) error {
	return nil
}

func (adp DummyAdapter) DataRequest(id string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"id": id,
	}, nil
}

func main() {
	r := mux.NewRouter()
	
	adp := DummyAdapter{}
	
	integrase.Prepare(adp, integrase.MuxWrap{Router: r})
	
	// Add any middleware here (ex: logging middleware)
	// Add logging middleware
	log := handlers.LoggingHandler(os.Stdout, r)

	http.ListenAndServe(":8080", log)
}
