package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
	"log"
	"qin/configs/consts"
	"qin/web"
	"time"
)

const TokenExpireDuration = JWTOverTime

var MySecret = []byte(JWTSecret)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:" + consts.RedisPort,
		DB:   0, // use default DB
	})
}

func GenToken(userInfo *web.User) (string, error) {
	c := MyClaims{
		userInfo.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "kangning",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完成的编码后的字符串token
	tokenStr, err := token.SignedString(MySecret)
	if err != nil {
		log.Panicln(err)
	}
	//从黑名单中取出token,该想法不合适，已禁止
	//num, err1 := Client.HDel(context.Background(), "black", strconv.Itoa(int(userInfo.Id))).Result()
	//log.Println("Client.LRem return ", num, err1)
	return tokenStr, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i any, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("JWT 解析错误")
}
