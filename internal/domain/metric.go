package domain

// Metric represents telecom data
type Metric struct {
	TowerID    string  `json:"tower_id"`
	LatencyMs  int     `json:"latency_ms"`
	Users      int     `json:"users"`
	PacketLoss float64 `json:"packet_loss"`
	Timestamp  int64   `json:"timestamp"`
}
