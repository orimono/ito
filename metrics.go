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

type CPUMetrics struct {
	UsedPercent float64   `json:"used_percent"`
	PerCPU      []float64 `json:"per_cpu"`
}

type DiskPartition struct {
	Mountpoint  string  `json:"mountpoint"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

type DiskMetrics struct {
	Partitions []DiskPartition `json:"partitions"`
}

type NetworkInterface struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytes_sent"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
}

type NetworkMetrics struct {
	Interfaces []NetworkInterface `json:"interfaces"`
}

type LoadMetrics struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}
