package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func GenerateJwt(uid uint64) (string, *time.Time, error) {
	mySigningKey := []byte(viper.GetString("jwt.secret_key"))

	maxAge := viper.GetInt("jwt.max_age")     // read from configuration file
	expireTime := time.Now().Add(time.Duration(maxAge) * time.Second)

	claims := &jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    viper.GetString("jwt.issuer"),
		Subject:   strconv.Itoa(int(uid)),      // you may want to encrypt this
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		logrus.Error(err)
		return "", nil, err
	}
	return tokenString, &expireTime, nil
}

func ParseJwt(tokenString string) (*jwt.Token, error) {
	mySigningKey := []byte(viper.GetString("jwt.secret_key"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpexted singing method: %v\n", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
