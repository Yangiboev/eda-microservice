package catalog_service

type Brand struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	PreviewText string `json:"preview_text"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Order       int64  `json:"order"`
	Image       string `json:"image"`
}

type CreateBrand struct {
	Name        string `json:"name" binding:"required"`
	PreviewText string `json:"preview_text"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Order       int64  `json:"order"`
	Image       string `json:"image"`
}

type GetAllBrands struct {
	Page   int64  `json:"page"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}
