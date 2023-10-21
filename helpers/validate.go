package helpers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gaurav2721/userSrv/models"
	"github.com/go-playground/validator/v10"
)

func ConvertTimeToUnixMilliSecond(minTimeInYear int) (int64, error) {
	date := time.Date(minTimeInYear, 1, 1, 0, 0, 0, 0, time.UTC)
	return date.UnixMilli(), nil
}

func MinTimeAllowed(inField validator.FieldLevel) bool {

	value, kind, nullable := inField.ExtractType(inField.Field())
	param := inField.Param()

	log.Printf("=> name: [%s]\n", inField.FieldName())
	log.Printf("=> tag: [%s]\n", inField.GetTag())
	log.Printf("=> struct field name: [%s]\n", inField.StructFieldName())
	log.Printf("=> param: [%s]\n", param)
	log.Printf("=> value: [%s]\n", value.String())
	log.Printf("=> kind: [%s]\n", kind.String())
	if nullable {
		log.Printf("=> nullable: true\n")
	} else {
		log.Printf("=> nullable: false\n")
	}

	if kind.String() != "int64" {
		log.Printf("error invalid time passed ")
		return false
	}

	year, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return false
	}
	minTimeMilliSeconds, err := ConvertTimeToUnixMilliSecond(int(year))
	if err != nil {
		return false
	}
	if value.Int() > minTimeMilliSeconds {
		return true
	}

	return false
}

func ValidateUser(user models.User) error {

	myValidator := validator.New()
	if err := myValidator.RegisterValidation("MinTimeAllowed", MinTimeAllowed); err != nil {
		panic(err)
	}

	err := myValidator.Struct(user)

	if err != nil {
		msg := fmt.Sprintf("error validating the input user : %s", err.Error())
		log.Println(msg)
		return errors.New(msg)
	} else {
		return nil
	}
}
