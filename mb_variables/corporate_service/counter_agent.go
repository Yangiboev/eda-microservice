package corporate_service

type LegalCounterAgent struct {
	ID            string `json:"id"`
	Name          string `json:"name" binding:"required"`
	Telephone     string `json:"telephone" binding:"required"`
	Inn           int32 `json:"inn" binding:"required"`
	Mfo           int32 `json:"mfo" binding:"required"`
	AccountNumber int64 `json:"account_number" binding:"required"`
	Address       string `json:"address"  binding:"required"`
	Description   string `json:"description"`
}

type PhysicalCounterAgent struct {
	ID            string `json:"id"`
	Name          string `json:"name" binding:"required"`
	Telephone     string `json:"telephone" binding:"required"`
	Address       string `json:"address"  binding:"required"`
	Description   string `json:"description"`
}

type CreateLegalCounterAgent struct {
	Name          string `json:"name" binding:"required"`
	Telephone     string `json:"telephone" binding:"required"`
	Inn           int32 `json:"inn" binding:"required"`
	Mfo           int32 `json:"mfo" binding:"required"`
	AccountNumber int64 `json:"account_number" binding:"required"`
	Address       string `json:"address"  binding:"required"`
	Description   string `json:"description"`
}

type CreatePhysicalCounterAgent struct {
	Name          string `json:"name" binding:"required"`
	Telephone     string `json:"telephone" binding:"required"`
	Address       string `json:"address"  binding:"required"`
	Description   string `json:"description"`
}
type GetAllCounterAgents struct {
	Page        int64  `json:"page"`
	Limit       int64  `json:"limit"`
	Name 		string `json:"name"`
}

