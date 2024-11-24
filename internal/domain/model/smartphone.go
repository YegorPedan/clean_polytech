package model

import "time"

type Smartphone struct {
	ID             string
	Model          string
	Charge         int
	ConnectionTime time.Time
	DisconnectTime time.Time
	UserID         string
}
