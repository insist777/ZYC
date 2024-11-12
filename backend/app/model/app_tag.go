package model

type AppTag struct {
	BaseModel
	AppId uint `json:"appId" gorm:"type:bigint;not null"`
	TagId uint `json:"tagId" gorm:"type:int;not null"`
}
