package ingestion

import (
	"encoding/json"
	"fmt"
	"net/http"
	"telecom_anomaly_engine/internal/domain"
)

func MetriHandler(w http.ResponseWriter, r *http.Request) {

	var metric domain.Metric

	//Decodes the coming request into Metric structure
	err := json.NewDecoder(r.Body).Decode(&metric)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
	}

	if metric.TowerID == "" {
		http.Error(w, "towerID missing", http.StatusBadRequest)
	}

	//Print the decoded message
	fmt.Printf("Received: %+v\n", metric)

	//Send success Response
	w.WriteHeader(http.StatusOK)

}
