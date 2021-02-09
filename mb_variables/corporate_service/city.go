package corporate_service
type City struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
}

type CreateCity struct {
	Name        string `json:"name" binding:"required"`
}

