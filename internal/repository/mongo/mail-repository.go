package mongo_repository

import (
	"context"
	"github/eugene-krivtsov/idler-email/internal/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type MailRepository struct {
	db *mongo.Collection
}

func NewMailRepository(db *mongo.Database) *MailRepository {
	return &MailRepository{
		db: db.Collection(mailCollection),
	}
}

func (r *MailRepository) Save(ctx context.Context, message entity.Mail) error {
	_, err := r.db.InsertOne(ctx, message)
	return err
}
