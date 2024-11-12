package model

type AppInstallResource struct {
	BaseModel
	AppInstallId uint   `json:"appInstallId" gorm:"type:bigint;not null;"`
	LinkId       uint   `json:"linkId"  gorm:"type:int;not null;"`
	ResourceId   uint   `json:"resourceId" gorm:"type:bigint;"`
	Key          string `json:"key" gorm:"type:varchar(64);not null"`
	From         string `json:"from" gorm:"type:varchar(64);not null;default:local"`
}
