package passwords

import "GoPasswordManager/helpers"

func AddEntry() helpers.Entry {
	println("STARTING ADDING")
	email := helpers.GetStrInput("Email: ")
	username := helpers.GetStrInput("Username: ")
	password := helpers.GetStrInput("Password: ")
	println(email, username, password, "Saved")
	entry := helpers.Entry{
		Email:    email,
		Username: username,
		Password: password,
		Website:  "",
		Category: "",
		Note:     "",
	}

	return entry
}
