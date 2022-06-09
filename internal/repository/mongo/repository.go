package mongo

import (
	"context"
	"github/eugene-krivtsov/idler-email/internal/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	mailCollection = "email"
)

type Mails interface {
	Save(ctx context.Context, message entity.Mail) error
}

type Repositories struct {
	Mails Mails
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Mails: NewMailRepository(db),
	}
}
