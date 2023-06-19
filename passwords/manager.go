package passwords

import (
	"GoPasswordManager/helpers"
	"fmt"
)

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

func ViewEntries() {
	var entries = helpers.LoadEntries("./passwords.json")
	for i, v := range entries {
		fmt.Println(i, v.Website, v.Username)
	}
}

func DeleteEntry(entries []helpers.Entry) []helpers.Entry {
	//Remove element at index from entries
	index := helpers.GetIntInput()
	return append(entries[:index], entries[index+1:]...)
}
