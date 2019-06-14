package model

type Permission struct {
	tableName struct{} `sql:"permission"`

	ID          int32  `json:"id,omitempty"`
	Application string `json:"application,omitempty"`
	Authority   string `json:"authority,omitempty"`
	Description string `json:"description,omitempty"`
}
