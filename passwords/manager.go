package passwords

import (
	"GoPasswordManager/helpers"
	"errors"
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

func ViewEntries() error {
	var entries = helpers.LoadEntries("./passwords.json")
	if len(entries) == 0 {
		return errors.New("no entries found")
	}
	for i, v := range entries {
		fmt.Println(i, v.Website, v.Username)
	}
	return nil
}

func DeleteEntry(entries []helpers.Entry) []helpers.Entry {
	//Remove element at index from entries
	index := helpers.GetIntInput()
	return append(entries[:index], entries[index+1:]...)
}

func AreYouSure() bool {
	//ask the user for confirmation, write the code for me
	resp := helpers.GetStrInput("Are you sure? (y/n): ")
	if resp == "y" || resp == "Y" || resp == "yes" || resp == "YES" {
		return true
	}

	return false

}
