package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (app *App) dayHighlightsAPIHandler(w http.ResponseWriter, r *http.Request) {
	day := r.FormValue("day")

	timestamp, err := time.Parse("2006-01-02", day)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	highlights, err := app.db.GetDayHighlights(timestamp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(highlights)
}

func (app *App) highlightAPIHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	highlightID := vars["id"]
	highlight, err := app.db.GetHighlight(highlightID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(highlight)
}

func (app *App) highlightMirrorsAPIHandler(w http.ResponseWriter, r *http.Request) {
	highlightID := r.FormValue("highlightId")
	highlightMirrors, err := app.db.GetHighlightMirrors(highlightID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(highlightMirrors)
}
