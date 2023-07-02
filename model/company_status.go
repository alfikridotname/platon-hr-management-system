package model

type CompanyStatus struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (CompanyStatus) TableName() string {
	return "company_status"
}
