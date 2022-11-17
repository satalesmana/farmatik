package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenMetadata ...
type TokenMetadata struct {
	ID    string
	Name  string
	Email string
}

// CreateToken ...
func CreateToken(id, name, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["name"] = name
	claims["email"] = email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// TokenValid ...
func TokenValid(headerToken string) error {
	token, err := verifyToken(headerToken)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata ...
func ExtractTokenMetadata(headerToken string) (*TokenMetadata, error) {
	token, err := verifyToken(headerToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// id, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		// if err != nil {
		// 	return nil, err
		// }

		id, ok := claims["id"].(string)
		if !ok {
			return nil, err
		}

		name, ok := claims["name"].(string)
		if !ok {
			return nil, err
		}

		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}

		return &TokenMetadata{
			ID:    id,
			Name:  name,
			Email: email,
		}, nil
	}
	return nil, err
}

func verifyToken(headerToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
