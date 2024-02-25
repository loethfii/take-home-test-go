package helper

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(p []byte) string {
	psw, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	
	return string(psw)
}

func ComparePassword(hashedPassword, psw []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, psw)
	return err == nil
}
