package jwt

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// iss (issuer)：签发人
// sub (subject)：主题
// aud (audience)：受众
// exp (expiration time)：过期时间
// nbf (Not Before)：生效时间，在此之前是无效的
// iat (Issued At)：签发时间
// jti (JWT ID)：编号

type Jwt struct {
	Issuer       string `yaml:"Issuer"`
	Subject      string `yaml:"Subject"`
	ExpireSecond int64  `yaml:"ExpireSecond"`

	SignKey       string            `yaml:"sign_key"`
	signKey       []byte            `yaml:"-"`
	SigningMethod string            `yaml:"signing_method"`
	signingMethod jwt.SigningMethod `yaml:"-"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (p *Jwt) Init() error {
	p.signingMethod = jwt.GetSigningMethod(p.SigningMethod)
	if p.signingMethod == nil {
		return fmt.Errorf("jwt not support SigningMethod: %s", p.SigningMethod)
	}

	if p.ExpireSecond <= 0 {
		p.ExpireSecond = 10
	}

	p.signKey = []byte(p.SignKey)

	return nil
}

func (p *Jwt) CreateToken(user, ip string) (string, error) {
	now := time.Now().Unix()

	claims := jwt.StandardClaims{
		Audience:  ip,
		ExpiresAt: now + p.ExpireSecond,
		Id:        user,
		IssuedAt:  now,
		Issuer:    p.Issuer,
		Subject:   p.Subject,
		NotBefore: now,
	}

	return jwt.NewWithClaims(p.signingMethod, claims).SignedString(p.signKey)
}

func (p *Jwt) Verify(tokenString, ip string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			return nil, errors.New("unexpected Claims Type")
		}
		if claims.Issuer != p.Issuer {
			return nil, errors.New("unexpected Issuer")
		}
		if claims.Subject != p.Subject {
			return nil, errors.New("unexpected Subject")
		}
		if e = claims.Valid(); e != nil {
			return
		}

		return p.signKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("token parse failed, err: %v", err)
	}
	if token == nil {
		return "", errors.New("token parse failed, token is nil")
	}

	return token.Claims.(*jwt.StandardClaims).Id, nil
}
