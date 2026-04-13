package ito

import "encoding/json"

// ── Risk ─────────────────────────────────────────────────────────────────────

type RiskLevel int

const (
	RiskLow      RiskLevel = iota // 只读或无副作用
	RiskMedium                    // 可逆变更，如服务重启
	RiskHigh                      // 难以撤销，如关机、配置修改
	RiskCritical                  // 破坏性操作，如格式化、删除
)

// ── Capability ────────────────────────────────────────────────────────────────

type Capability struct {
	Kind      string   `json:"kind"`
	Version   string   `json:"version"`
	Platforms []string `json:"platforms"`

	Risk              RiskLevel `json:"risk"`
	RequiresElevation bool      `json:"requires_elevation"`
	Warning           string    `json:"warning,omitempty"`
}

// ── Envelope (WebSocket wire format between loom and shutter) ─────────────────

type MessageKind string

const (
	KindJoinAccepted       MessageKind = "join.accepted"
	KindTaskRequest        MessageKind = "task.request"
	KindTaskResult         MessageKind = "task.result"
	KindExecutorRegister   MessageKind = "executor.register"
	KindExecutorRegistered MessageKind = "executor.registered"
)

type Envelope struct {
	Kind    MessageKind     `json:"kind"`
	Payload json.RawMessage `json:"payload"`
}

func Encode(kind MessageKind, payload any) ([]byte, error) {
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return json.Marshal(Envelope{Kind: kind, Payload: p})
}

func Decode(data []byte) (*Envelope, error) {
	var env Envelope
	if err := json.Unmarshal(data, &env); err != nil {
		return nil, err
	}
	return &env, nil
}

// ── Task ──────────────────────────────────────────────────────────────────────

type TaskRequest struct {
	TaskID string          `json:"task_id"`
	Kind   string          `json:"kind"`
	Params json.RawMessage `json:"params"`
}

type TaskResult struct {
	TaskID  string          `json:"task_id"`
	Success bool            `json:"success"`
	Output  json.RawMessage `json:"output,omitempty"`
	Error   string          `json:"error,omitempty"`
}

// ── Executor registration ─────────────────────────────────────────────────────

type ExecutorRegistration struct {
	Kind              string    `json:"kind"`
	Version           string    `json:"version"`
	Runtime           string    `json:"runtime"` // bash | python | powershell
	Script            string    `json:"script"`
	Platforms         []string  `json:"platforms"`
	Risk              RiskLevel `json:"risk"`
	RequiresElevation bool      `json:"requires_elevation"`
	Warning           string    `json:"warning,omitempty"`
}

// ── Join ──────────────────────────────────────────────────────────────────────

type JoinPacket struct {
	NodeID   string `json:"node_id"`
	Hostname string `json:"hostname"`

	OS   string `json:"os"`
	Arch string `json:"arch"`

	TaskManifest map[string]Capability `json:"task_manifest"`

	Version string   `json:"version"`
	Tags    []string `json:"tags"`
}
