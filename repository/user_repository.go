package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepository(coll *mongo.Collection) *UserRepository {
	return &UserRepository{col: coll}

}

func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = now
	u.IsActive = true
	if u.Role == "" {
		u.Role = "customer"
	}
	_, err := r.col.InsertOne(ctx, u)
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}
