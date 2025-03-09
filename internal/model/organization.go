package model

type Organization struct {
	OrganizationID int `gorm:"primaryKey;autoIncrement" json:"organization_id"`
}
