package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	auth "github.com/roronoazor/InventoryBackendQL/auth/models"
	"github.com/roronoazor/InventoryBackendQL/database"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetUserByEmail(c *fiber.Ctx, email string) (*auth.User, error) {
	var user auth.User
	userCollection := database.MongoDbInstance.Db.Collection("users")
	if err := userCollection.FindOne(c.Context(), bson.D{{Key: "email", Value: email}}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(c *fiber.Ctx) (*auth.User, error) {

	userCollection := database.MongoDbInstance.Db.Collection("users")

	// create a new user struct
	// Parse body into struct
	user := new(auth.User)
	if err := c.BodyParser(user); err != nil {
		return nil, err
	}

	user.ID = ""

	// validate if a user with that email already exists
	if existingUser, err := GetUserByEmail(c, user.Email); err != nil {
		return nil, err
	} else if existingUser != nil {
		return nil, errors.New("User with this email already exists")
	}

	// add record
	insertionResult, err := userCollection.InsertOne(c.Context(), user)
	if err != nil {
		return nil, err
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := userCollection.FindOne(c.Context(), filter)

	// decode the Mongo record into Employee
	createdUser := new(auth.User)
	createdRecord.Decode(createdUser)

	return createdUser, nil
}

func authentiateUser(c *fiber.Ctx, email string, password string) (*auth.User, error) {
	userPointer, err := GetUserByEmail(c, email)
	if err != nil {
		return nil, err
	}
	// check passwords
	hashed, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	// raise invalid credentials error

}
