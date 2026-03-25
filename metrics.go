package ito

type Telemetry struct {
	NodeID    string `json:"node_id"`
	Timestamp int64  `json:"ts"`
	Type      string `json:"type"`
	Payload   any    `json:"payload"`
}

type MemoryMetrics struct {
}
