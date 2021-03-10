package utils

// Example HTTP auth using asymmetric crypto/RSA keys
// This is based on a (now outdated) example at https://gist.github.com/cryptix/45c33ecf0ae54828e63b

import (
	"apitest/models"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// location of the files used for signing and verification
const (
	privKeyPath = "./utils/private.key" // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "./utils/public.key"  // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey   *rsa.PublicKey
	signKey     *rsa.PrivateKey
	verifyBytes []byte
	secretKey   string
	verpk       bool
)

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err = ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)

	secretKey = "0816477729"
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Define some custom types were going to use within our tokens
type CustomerInfo struct {
	Username string         `json:"username"`
	Level    uint           `json:"level"`
	ID       uint           `json:"id"`
	Roles    []*models.Role `json:"roles"`
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	CustomerInfo
}

func JwtSign(payload models.User) string {
	return jwtpkSign(payload)
	// jwtSecretSign(payload)
}

func JwtVerify(c *gin.Context) {
	jwtPubkVerify(c)
	// jwtSecretVerify(c)
}

func jwtpkSign(payload models.User) string {
	log.Println("-------private key sign jwt token---------", payload)
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
		CustomerInfo{payload.Name, payload.Level, payload.ID, payload.Roles},
	}
	token, err := t.SignedString(signKey)
	fatal(err)
	return token
}

func jwtPubkVerify(c *gin.Context) {
	log.Println("-----verify pubkey-----")
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
	}
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	log.Println("jwt->", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	fatal(err)
	log.Println("token-->", token.Claims.Valid())
	log.Println("tokenx-->", token.Claims)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println("claims-->", claims)
		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["username"])
		level := fmt.Sprintf("%v", claims["level"])
		c.Set("jwt_staff_id", staffID)
		c.Set("jwt_username", username)
		c.Set("jwt_level", level)
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "nok", "message": "invalid token", "error": err})
		c.Abort()
	}
}

func jwtSecretSign(payload models.User) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	// Payload end

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token
}

func jwtSecretVerify(c *gin.Context) {

	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
	}
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)

		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["username"])
		level := fmt.Sprintf("%v", claims["level"])
		c.Set("jwt_staff_id", staffID)
		c.Set("jwt_username", username)
		c.Set("jwt_level", level)

		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "nok", "message": "invalid token", "error": err})
		c.Abort()
	}
}

func GetPubkey() string {
	pubk := string(verifyBytes)
	if pubk != "" {
		return string(verifyBytes)
	} else {
		return "---publick key---"
	}

}
