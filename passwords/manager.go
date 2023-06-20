package passwords

import (
	"GoPasswordManager/helpers"
	"errors"
	"fmt"
)

func AddEntry(masterPassword string) helpers.Entry {
	println("STARTING ADDING")
	email, _ := EncryptPassword(helpers.GetStrInput("Email: "), masterPassword)
	username, _ := EncryptPassword(helpers.GetStrInput("Username: "), masterPassword)
	password, _ := EncryptPassword(helpers.GetStrInput("Password: "), masterPassword)
	website, _ := EncryptPassword(helpers.GetStrInput("Website: "), masterPassword)
	category, _ := EncryptPassword(helpers.GetStrInput("Category: "), masterPassword)
	notes, _ := EncryptPassword(helpers.GetStrInput("Note: "), masterPassword)

	entry := helpers.Entry{
		Email:    email,
		Username: username,
		Password: password,
		Website:  website,
		Category: category,
		Note:     notes,
	}

	return entry
}

func ViewEntries(masterPassword string) ([]helpers.Entry, error) {
	var entries = helpers.LoadEntries("./passwords.json")
	if len(entries) == 0 {
		fmt.Println("No entries found")
		entries := []helpers.Entry{helpers.Entry{}}
		return entries, errors.New("no entries found")
	}
	for i, v := range entries {
		decryptedemail, _ := DecryptPassword(v.Email, masterPassword)
		fmt.Println(i, "Entry: ", decryptedemail)
	}
	return entries, nil
}

func ViewEntry(entry helpers.Entry, masterPassword string) {
	email, _ := DecryptPassword(entry.Email, masterPassword)
	username, _ := DecryptPassword(entry.Username, masterPassword)
	password, _ := DecryptPassword(entry.Password, masterPassword)
	website, _ := DecryptPassword(entry.Website, masterPassword)
	category, _ := DecryptPassword(entry.Category, masterPassword)
	notes, _ := DecryptPassword(entry.Note, masterPassword)
	fmt.Println("Email: ", email)
	fmt.Println("Username: ", username)
	fmt.Println("password:", password)
	fmt.Println("website:", website)
	fmt.Println("category:", category)
	fmt.Println("notes:", notes)
}

func DeleteEntry(entries []helpers.Entry) []helpers.Entry {
	//Remove element at index from entries
	index := helpers.GetIntInput()
	return append(entries[:index], entries[index+1:]...)
}

func AreYouSure() bool {
	//ask the user for confirmation, write the code for me
	resp := helpers.GetStrInput("Are you sure? (y/n): ")
	if resp == "y" || resp == "Y" || resp == "yes" || resp == "YES" || resp == "1" {
		return true
	}

	return false
}
