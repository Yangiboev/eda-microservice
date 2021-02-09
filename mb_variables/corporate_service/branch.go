package corporate_service

type Branch struct {
	ID                string `json:"id"`
	CompanyId         string `json:"company_id" binding:"required"`
	CityId            string `json:"city_id" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Telephone         string `json:"telephone" binding:"required"`
	NumberOfEmployees int64  `json:"account_number" binding:"required"`
	Address           string `json:"address"  binding:"required"`
	Description       string `json:"description"`
}

type CreateBranch struct {
	CompanyId         string `json:"company_id" binding:"required"`
	CityId            string `json:"city_id" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Telephone         string `json:"telephone" binding:"required"`
	NumberOfEmployees int64  `json:"account_number" binding:"required"`
	Address           string `json:"address"  binding:"required"`
	Description       string `json:"description"`
}

type GetAllBranches struct {
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
	Name  string `json:"name"`
}
