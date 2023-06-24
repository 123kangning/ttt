package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JWTSecret   = "kangning"
	JWTOverTime = time.Hour * 720
)

type MyClaims struct {
	UserName string
	jwt.StandardClaims
}
