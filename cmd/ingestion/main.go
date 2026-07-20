package main

import (
	"fmt"
	"log"
	"net/http"
	"telecom_anomaly_engine/internal/ingestion"
	"telecom_anomaly_engine/internal/streaming"
)

func main() {
	//new route declaration
	mux := http.NewServeMux()

	//mapping the route
	mux.HandleFunc("/metrics", ingestion.MetriHandler)

	/*Start worker*/
	for i := 1; i <= 5; i++ {
		go streaming.StartWorker(i)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Ingestion service running on http://localhost:8080")

	//server starts here
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
