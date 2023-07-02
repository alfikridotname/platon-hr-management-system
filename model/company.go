package model

type Company struct {
	ID            int           `gorm:"primary_key" json:"id"`
	Name          string        `json:"name"`
	StatusID      int           `json:"status_id"`
	CompanyStatus CompanyStatus `gorm:"foreignkey:StatusID" json:"company_status"`
}
