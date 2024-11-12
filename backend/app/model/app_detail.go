package model

type AppDetail struct {
	BaseModel
	AppId               uint   `json:"appId" gorm:"type:bigint;not null"`
	Version             string `json:"version" gorm:"type:varchar(64);not null"`
	Params              string `json:"-" gorm:"type:longtext;"`
	DockerCompose       string `json:"dockerCompose"  gorm:"type:longtext;"`
	Status              string `json:"status" gorm:"type:varchar(64);not null"`
	LastVersion         string `json:"lastVersion" gorm:"type:varchar(64);"`
	LastModified        int    `json:"lastModified" gorm:"type:bigint;"`
	DownloadUrl         string `json:"downloadUrl"  gorm:"type:varchar(255);"`
	DownloadCallBackUrl string `json:"downloadCallBackUrl" gorm:"type:longtext;"`
	Update              bool   `json:"update"`
	IgnoreUpgrade       bool   `json:"ignoreUpgrade"`
}
