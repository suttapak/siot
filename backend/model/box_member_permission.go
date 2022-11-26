package model

type BoxMemberPermission struct {
	Model
	// attribute
	CanRead     bool `json:"canRead"`
	CanWrite    bool `json:"canWrite"`
	BoxMemberId uint `json:"boxMemberId"`
}
