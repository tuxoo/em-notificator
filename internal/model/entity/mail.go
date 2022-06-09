package entity

import "time"

type Mail struct {
	Address string    `bson:"address"`
	Subject string    `bson:"subject"`
	SentAt  time.Time `bson:"sentAt"`
}
