package model

import "time"

type Smartphone struct {
	ID             int
	Model          string
	Charge         int
	ConnectionTime time.Time
	DisconnectTime time.Time
}
