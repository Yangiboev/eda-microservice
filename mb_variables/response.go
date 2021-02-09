package mb_variables

type Response struct {
	ID    string `json:"id"`
	Error Error  `json:"error"`
}
