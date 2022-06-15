package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type TokenInfo struct {
	Value string
	Expires time.Time
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

func setExpiresIn(days time.Duration) time.Time {
	day := 24 * time.Hour
	expirationTime := time.Now().Add(day * days)

	return expirationTime
}

func IssueJWT(uid uuid.UUID) (TokenInfo, error) {
	// Set expiration time, here defaulted to 30days
	expirationTime := setExpiresIn(30)

	// Create claims that will be in jwt payload
	claims := &Claims{
		UserID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	// Declare token with algorithm for signin + claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create JWT String
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return TokenInfo{}, err
	}

	// Return token info struct
	return TokenInfo{ Value: tokenStr, Expires: expirationTime}, err
}

func RenewJWT(tknStr string) (TokenInfo, error) {
	claims, err := CheckJWT(tknStr)
	if err != nil {
		return TokenInfo{}, err
	}

	// Check if claims expires in more than 1 day
	if claims.ExpiresAt.After(time.Now().Add(time.Hour * 24)) {
		return TokenInfo{}, errors.New("current token is still valid for a day")
	}

	expirationTime := setExpiresIn(30)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return TokenInfo{}, err
	}

	return TokenInfo{ Value: tokenStr, Expires: expirationTime }, err
}

func CheckJWT(tknStr string) (Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tknStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return *claims, err
	}

	if !token.Valid {
		return *claims, jwt.ErrSignatureInvalid
	}

	return *claims, err
}