package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/GoExpertCurso/auction-go/config/logger"
	"github.com/GoExpertCurso/auction-go/internal/entity/user_entity"
	"github.com/GoExpertCurso/auction-go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntitytMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func newUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{Collection: database.Collection("users")}
}

func (ur *UserRepository) FindUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntitytMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		logger.Error(fmt.Sprintf("User not found with this id = %d", userId), err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, internal_error.NewNotFoundError(fmt.Sprintf("User not found with this id = %d", userId))
		}

		logger.Error(fmt.Sprintf("Error tryung to find user by userId"), err)
		return nil, internal_error.NewInternalServerError("Error tryung to find user by userId")
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
