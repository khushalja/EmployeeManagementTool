package helper

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	EmployeeName string
	// EmployeeId   string
	// EmployeeType string
	jwt.RegisteredClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateAllTokens(EmployeeName string) (signedToken string, signedRefreshToken string, err error) {
	fmt.Println("line 21 in GenerateAlltoken", EmployeeName)
	claims := &SignedDetails{
		EmployeeName: EmployeeName,
		//converts a int to string if we only use string then the output will give rune and not digit string
		// EmployeeId:   string(rune(EmployeeId)),
		// EmployeeType: string(rune(EmployeeType)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}
	fmt.Println(claims.RegisteredClaims)
	fmt.Println(claims.EmployeeName)
	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(168))),
		},
	}
	var returnErr error
	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		returnErr = err
	}
	signedRefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		returnErr = err
	}
	fmt.Println(signedToken)
	return signedToken, signedRefreshToken, returnErr
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "token is invalid" + err.Error()
		// msg = err.Error()
		return
	}

	return claims, msg
}
