// Why: events define durable domain messages that other services can consume from Kafka.
// What to do: add event types here when the system needs to broadcast business facts.
package event

import "time"

const EventUserRegistered = "user.registered"

type UserRegisteredEvent struct {
	EventType string    `json:"event_type"`
	UserID    uint      `json:"user_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
