package helpers

import (
	"github.com/sethvargo/go-password/password"
)

func GeneratePassword() string {
	res, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		Danger(err.Error())
	}

	return res
}
