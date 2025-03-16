package model

type OrganizationModel struct {
	ID       uint   `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name;type:VARCHAR(255);not null;type:varchar(255)" json:"name" binding:"required"`
	StatusID uint   `gorm:"column:status_id;type:INTEGER;not null" json:"status_id" binding:"required"`
}

/*func (OrganizationModel) TableName() string {
	return "organization"
}*/

/*
type OrganizationStatusModel struct {
	ID   uint   `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:VARCHAR(20);not null" json:"name" binding:"required"`
}

func (OrganizationStatusModel) TableName() string {
	return "organization_statuses"
}
*/
