package ito

type Telemetry struct {
	NodeID    string `json:"node_id"`
	Timestamp int64  `json:"ts"`
	Type      string `json:"type"`
	Payload   any    `json:"payload"`
}

type MemoryMetrics struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}
