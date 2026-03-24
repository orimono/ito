package ito

type JoinPacket struct {
	NodeID   string `json:"node_id"`
	Hostname string `json:"hostname"`

	OS   string `json:"os"`
	Arch string `json:"arch"`

	TaskManifest map[string]string `json:"task_manifest"`

	Version string   `json:"version"`
	Tags    []string `json:"tags"`
}
