package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"telecom_anomaly_engine/internal/domain"
	"time"
)

func generateNormalMetric(tower_id string) domain.Metric {
	return domain.Metric{
		TowerID:    tower_id,
		LatencyMs:  rand.Intn(40) + 40,   //40-80 ms
		Users:      rand.Intn(200) + 200, //200-400 users
		PacketLoss: rand.Float64() * 2,   //2-4%
		Timestamp:  time.Now().UnixMilli(),
	}
}
func generateAbnormalMetric(tower_id string) domain.Metric {
	return domain.Metric{
		TowerID:    tower_id,
		LatencyMs:  rand.Intn(100) + 100, //100-200 ms
		Users:      rand.Intn(10),        //sudden drop users
		PacketLoss: rand.Float64()*5 + 5, //5-10%%
		Timestamp:  time.Now().UnixMilli(),
	}
}

func sendMetric(metric domain.Metric) {
	jsonByte, err := json.Marshal(metric)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	buf := bytes.NewBuffer(jsonByte)
	resp, err := http.Post(
		"http://localhost:8080/metrics",
		"application/json",
		buf,
	)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()
}

func simulateTower(tower_id string) {
	fmt.Println("simulation tower started:", tower_id)
	for {
		var metric domain.Metric
		if rand.Float64() < 0.1 {
			metric = generateAbnormalMetric(tower_id)
		} else {
			metric = generateNormalMetric(tower_id)
		}
		log.Printf("Simulated: %+v\n", metric)

		sendMetric(metric)

		time.Sleep(time.Duration(rand.Intn(50)+50) * time.Millisecond)
	}
}

func main() {
	fmt.Println("Telecom simulator starts")
	totalTower := 10

	for i := 0; i < totalTower; i++ {
		towerID := fmt.Sprintf("Tower-%d", i)
		go simulateTower(towerID)
	}
	for {
		time.Sleep(time.Second)
	}
	//select {}

}
