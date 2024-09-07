package responses

type body struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
