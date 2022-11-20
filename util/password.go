package util

import (

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string,error){
	hashbyte , err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(hashbyte), nil
}

func CheckPasswordAndHash(hashedPassword string,password string)error{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	return err
}