// Why: health DTOs keep the health response shape explicit and reusable.
// What to do: add only small status metadata here for probes and debugging.
package dto

type HealthResponse struct {
	App    string `json:"app"`
	Env    string `json:"env"`
	Status string `json:"status"`
}
