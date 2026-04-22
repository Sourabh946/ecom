package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("~ABC:123")

type Claims struct {
	User_id  int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, username string) (string, string, error) {

	// -----------------------------
	// ACCESS TOKEN (short-lived)
	// -----------------------------

	// Set expiration time (e.g., 15 minutes)
	accessExp := time.Now().Add(600 * time.Minute)

	// Create claims (payload) for access token
	accessClaims := &Claims{
		User_id:  userID,   // user identifier
		Username: username, // username (optional but useful)
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp), // expiry time
			Issuer:    "my-app",                      // identifies who issued the token
		},
	}

	// Create JWT object with claims using HMAC SHA256 signing method
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// Sign the token with secret key
	accessTokenStr, err := accessToken.SignedString(jwtKey)
	if err != nil {
		// return error if signing fails
		return "", "", err
	}

	// -----------------------------
	// REFRESH TOKEN (long-lived)
	// -----------------------------

	// Set longer expiration (e.g., 7 days)
	refreshExp := time.Now().Add(7 * 24 * time.Hour)

	// Create claims for refresh token
	// Note: Keep it minimal for better security
	refreshClaims := &Claims{
		User_id: userID, // only essential data
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp), // expiry time
			Issuer:    "my-app",
		},
	}

	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Sign refresh token
	refreshTokenStr, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Return both tokens
	return accessTokenStr, refreshTokenStr, nil
}

func VerifyJWT(tokenStr string) (*Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {

		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return jwtKey, nil
	})

	// 🔴 Handle errors properly
	if err != nil {

		// Token is expired
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired")
		}

		// Token is malformed / invalid
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf("malformed token")
		}

		// Signature invalid
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, fmt.Errorf("invalid signature")
		}

		return nil, fmt.Errorf("invalid token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
