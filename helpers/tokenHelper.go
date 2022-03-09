package helper

import "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	UserType  string
	jwt.StandardClaims
}
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")