package ito

type Capability struct {
	Version   string            `json:"version"`
	Params    map[string]string `json:"params"`
	Platforms []string          `json:"platforms"`
}

type JoinPacket struct {
	NodeID   string `json:"node_id"`
	Hostname string `json:"hostname"`

	OS   string `json:"os"`
	Arch string `json:"arch"`

	TaskManifest map[string]Capability `json:"task_manifest"`

	Version string   `json:"version"`
	Tags    []string `json:"tags"`
}
