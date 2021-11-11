package mongo

import (
	"errors"
	"log"
	"time"

	"github.com/vagnerlg/supermaketlist/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type MongoUser struct{}

func (r MongoUser) Insert(u domain.User) domain.User {

	MongoUser := user{
		Id:        primitive.NewObjectID(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := collection("user").InsertOne(ctx, MongoUser)

	if err != nil {
		log.Fatal(err)
		panic("Erro ao inserir novo item")
	}

	var oid primitive.ObjectID = result.InsertedID.(primitive.ObjectID)

	u.Id = oid.Hex()

	return u
}

func (r MongoUser) First(id string) (domain.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	u := user{}
	err = collection("user").FindOne(ctx, bson.D{{"_id", oid}}).Decode(&u)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	return domain.User{
		Id:       u.Id.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r MongoUser) FindByEmail(email string) (domain.User, error) {

	u := user{}
	err := collection("user").FindOne(ctx, bson.D{{"email", email}}).Decode(&u)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	return domain.User{
		Id:       u.Id.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r MongoUser) Login(email string, password string) bool {
	user, err := r.FindByEmail(email)
	if err != nil {
		return false
	}

	return user.ComparePassword(password)
}
