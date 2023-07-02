package response

import "hr-management-system/model"

type CompanyResponse struct {
	CompanyID         int    `json:"company_id"`
	CompanyName       string `json:"company_name"`
	CompanyStatusCode int    `json:"company_status_code"`
	CompanyStatusName string `json:"company_status_name"`
}

func ResponseCompany(company model.Company) CompanyResponse {
	return CompanyResponse{
		CompanyID:         company.ID,
		CompanyName:       company.Name,
		CompanyStatusCode: company.StatusID,
		CompanyStatusName: company.CompanyStatus.Name,
	}
}
