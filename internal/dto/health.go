package dto

type HealthResponse struct {
	App    string `json:"app"`
	Env    string `json:"env"`
	Status string `json:"status"`
}
