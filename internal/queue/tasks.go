// Why: queue tasks define the names and payloads shared by API producers and worker consumers.
// What to do: add one task type and payload struct here for each background job.
package queue

const TaskSendWelcomeEmail = "email:send_welcome"

type WelcomeEmailPayload struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}
