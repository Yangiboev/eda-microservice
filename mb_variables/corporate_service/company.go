package corporate_service

type Company struct {
	ID            string `json:"id"`
	Name          string `json:"name" binding:"required"`
	Telephone     string `json:"telephone" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Inn           int32 `json:"inn" binding:"required"`
	Mfo           int32 `json:"mfo" binding:"required"`
	AccountNumber int32 `json:"account_number" db:"account_number" binding:"required"`
	Address       string `json:"address"  binding:"required"`
	Description   string `json:"description"`
	Count 		  int32  `json:"-" binding:"required"`
}

type CreateCompany struct {
	Name          string `json:"name" binding:"required"`
	Telephone     string `json:"telephone" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Inn           int32 `json:"inn" binding:"required"`
	Mfo           int32 `json:"mfo" binding:"required"`
	AccountNumber int32 `json:"account_number" binding:"required"`
	Address       string `json:"address"  binding:"required"`
	Description   string `json:"description"`
}

type GetAllCompanies struct {
	Page       	int32  `json:"page"`
	Limit       int32  `json:"limit"`
	Name 		string `json:"name"`
}

type GetAllCompaniesResponse struct {
	Count       int32  `json:"count"`
	Companies 	[]*Company `json:"companies"`
}

type GetCompanyResponse struct {
	Company *Company `json:"company"`
}

