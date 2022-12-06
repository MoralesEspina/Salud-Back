package models

type Permission struct {
	Uuid           string `json:"uuid,omitempty"`
	Register       int    `json:"register,omitempty"`
	SubmittedAt    string `json:"submittedAt,omitempty"`
	ModifiedAt     string `json:"modifiedAt,omitempty"`
	PermissionDate string `json:"permissionDate,omitempty"`
	UuidPerson     string `json:"uuidPerson,omitempty"`
	BossOne        string `json:"bossOne,omitempty"`
	BossTwo        string `json:"bossTwo,omitempty"`
	Motive         string `json:"motive,omitempty"`
	StatusBossOne  string `json:"statusBossOne,omitempty"`
	StatusBossTwo  string `json:"statusBossTwo,omitempty"`
	Reason         string `json:"reason,omitempty"`
	Status         string `json:"status,omitempty"`
}
