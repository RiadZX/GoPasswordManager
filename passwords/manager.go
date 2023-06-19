package passwords

import "GoPasswordManager/helpers"

func AddEntry() helpers.Entry {
	println("STARTING ADDING")
	email := helpers.GetStrInput("Email: ")

	entry := helpers.Entry{
		Email:    email,
		Username: "username",
		Password: "password",
		Website:  "c",
		Category: "c",
		Note:     "c",
	}

	return entry
}
