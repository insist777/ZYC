package model

type Setting struct {
	BaseModel
	Key   string `json:"keys" gorm:"column:keys;type:varchar(255);not null;"`
	Value string `json:"value" gorm:"type:varchar(255)"`
	About string `json:"about" gorm:"type:longText"`
}
