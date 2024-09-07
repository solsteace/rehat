package apiResponses

type response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
