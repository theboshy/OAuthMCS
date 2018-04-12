package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

func EncrypthUni(userPassword string) (string,error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		//fmt.Print(err)
		return "errror",err
	}
	return string(hash),nil
}

func ValidateCrypthUni(providePassword string,hashFromDatabase string) (bool,error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashFromDatabase), []byte(providePassword)); err != nil {
		//fmt.Print(err)
		return false,err
	}
	return true,nil
}