package passwords

import "GoPasswordManager/helpers"

type Entry struct {
	email    string
	username string
	password string
	website  string
	category string
	note     string
}

func AddEntry() {
	println("STARTING ADDING")
	email := helpers.GetStrInput()
	username := helpers.GetStrInput()
	password := helpers.GetStrInput()
	println(email, username, password, "Saved")
	//entry := Entry{
	//	email:    email,
	//	username: username,
	//	password: password,
	//	website:  "",
	//	category: "",
	//	note:     "",
	//}

	helpers.SavePassword("Saved")
}
