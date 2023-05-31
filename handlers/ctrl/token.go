package ctrl

import (
	"fmt"
	"time"

	"github.com/CodingCookieRookie/audit-log/constants"
	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/golang-jwt/jwt"
	"gopkg.in/gomail.v2"
)

const (
	sender    = "log71536@gmail.com"
	smtpHost  = "smtp.gmail.com"
	smtpPort  = 587
	emailBody = "Hi thanks for choosing Canonicalvin audit-log service APIs.\nPlease use this as your API token:\n%v"
)

func SendJWTTokenToEmail(userEmail string) error {
	tokenString, err := createJWTTokenFromUserEmail(userEmail)

	if err != nil {
		log.Errorf("error signing token, err: %v", err)
		return err
	}

	if err := sendEmailToRecepient(tokenString, userEmail); err != nil {
		log.Errorf("error sending email, err: %v", err)
		return err
	}
	return nil
}

func createJWTTokenFromUserEmail(userEmail string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	claims["email"] = userEmail

	return token.SignedString([]byte(constants.GetAPIKey()))
}

func sendEmailToRecepient(tokenString, recepient string) error {
	password := constants.GetServiceEmailPassword()

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recepient)
	m.SetHeader("Subject", "API token email")
	m.SetBody("text/plain", fmt.Sprintf(emailBody, tokenString))

	d := gomail.NewDialer(smtpHost, smtpPort, sender, password)

	return d.DialAndSend(m)
}

func VerifyJWTToken(email, token string) error {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.GetAPIKey()), nil
	})
	if err != nil {
		log.Errorf("token parse err: %v", err)
		return err
	}

	if err = jwtToken.Claims.Valid(); err == nil {
		claims := jwtToken.Claims.(jwt.MapClaims)
		emailOfTokenRecepient := claims["email"]
		log.Infof("email of token recepient: %v", emailOfTokenRecepient)

		if emailOfTokenRecepient != email {
			log.Errorf("email of token recepient not the same as current user email, current user emai: %v", email)
			return fmt.Errorf("Please use a valid token recepient email to access this API service")
		}
	}

	return err
}
