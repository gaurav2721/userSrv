package helpers

import (
	"testing"

	"github.com/gaurav2721/userSrv/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserSuccess(t *testing.T) {

	//array of input for which ValidateUser should pass
	successInputs := []models.User{
		{
			Id:         "1",
			Name:       "Gaurav",
			SignupTime: 1234567,
		},
		{
			Id:         "2",
			Name:       "Saurav",
			SignupTime: -1234567,
		},
	}

	for _, input := range successInputs {
		err := ValidateUser(input)
		assert.Equal(t, nil, err)
	}
}

func TestValidateUserFailure(t *testing.T) {

	//array of input for which ValidateUser should fail
	failureInputs := []models.User{
		{
			Id:         "1",
			Name:       "G",
			SignupTime: 1234567,
		},
		{
			Id:         "2",
			Name:       "Saurav",
			SignupTime: -3690300000000000,
		},
		{
			Id:   "3",
			Name: "Saurav",
		},
		{
			Id:         "2",
			SignupTime: 1234567,
		},
		{
			Name:       "Saurav",
			SignupTime: 1234567,
		},
	}

	for _, input := range failureInputs {
		err := ValidateUser(input)
		t.Log("error", err)
		if err == nil {
			t.Fatal("input does not cause error", input)
		}
	}
}
