package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type ChartRequest struct {
	ChartType string      `json:"chartType"`
	Data      interface{} `json:"data"`
	Title     string      `json:"title"`
}

type ChartResponse struct {
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Options interface{} `json:"options"`
}

func parseJSONFeed(w http.ResponseWriter, r *http.Request) {
	var requestData interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Process the JSON feed (example: here we simply echo it back)
	response := ChartResponse{
		Type:    "line",
		Data:    requestData,
		Options: map[string]interface{}{"responsive": true},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func parseCSVFeed(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Error parsing CSV", http.StatusInternalServerError)
		return
	}

	// Process CSV data (example: convert to Chart.js format)
	labels := []string{}
	data := []int{}
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		labels = append(labels, record[0])
		var value int
		fmt.Sscanf(record[1], "%d", &value)
		data = append(data, value)
	}

	response := ChartResponse{
		Type: "bar",
		Data: map[string]interface{}{
			"labels": labels,
			"datasets": []map[string]interface{}{
				{
					"label":           "Dataset",
					"data":            data,
					"backgroundColor": "#36A2EB",
				},
			},
		},
		Options: map[string]interface{}{"responsive": true},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleChartAPI(w http.ResponseWriter, r *http.Request) {
	var chartRequest ChartRequest
	if err := json.NewDecoder(r.Body).Decode(&chartRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate Chart.js compatible JSON
	response := ChartResponse{
		Type:    chartRequest.ChartType,
		Data:    chartRequest.Data,
		Options: map[string]interface{}{"responsive": true},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	// Routes for feeds
	r.HandleFunc("/api/json-feed", parseJSONFeed).Methods("POST")
	r.HandleFunc("/api/csv-feed", parseCSVFeed).Methods("POST")

	// Route for chart API
	r.HandleFunc("/api/charts", handleChartAPI).Methods("POST")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}
