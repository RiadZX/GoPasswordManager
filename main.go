package main

import (
	"GoPasswordManager/helpers"
	"GoPasswordManager/passwords"
	"fmt"
	"os"
)

var (
	OPTIONS []string = []string{
		"New Entry",
		"Delete Entry",
		"View Entries",
		"Change master password",
		"Exit",
	}
)

func newEntry(masterPassword *string, logger *helpers.Logger) {
	entry := passwords.AddEntry(*masterPassword)
	helpers.SaveEntry(entry)
	logger.Log("Added Password Success")
}

func deleteEntry(masterPassword *string) {
	passwords.ViewEntries(*masterPassword)

	entries := passwords.DeleteEntry(helpers.LoadEntries("./passwords.json"))
	if passwords.AreYouSure() {
		helpers.SaveEntries(entries)
	}
}

func viewEntries(masterPassword *string) {
	entries, _ := passwords.ViewEntries(*masterPassword)
	entrynum := helpers.GetIntInput()
	passwords.ViewEntry(entries[entrynum], *masterPassword)
}

func changeMasterPassword(masterPassword *string) {

	newMasterPassword := helpers.GetStrInput("Enter new master password: ")
	//load current passwords, decrypt them, encrypt them with new key, and save them
	entries := helpers.LoadEntries("./passwords.json")
	var newentries []helpers.Entry
	for i, v := range entries {
		email, _ := passwords.DecryptPassword(v.Email, *masterPassword)
		username, _ := passwords.DecryptPassword(v.Username, *masterPassword)
		password, _ := passwords.DecryptPassword(v.Password, *masterPassword)
		website, _ := passwords.DecryptPassword(v.Website, *masterPassword)
		category, _ := passwords.DecryptPassword(v.Category, *masterPassword)
		notes, _ := passwords.DecryptPassword(v.Note, *masterPassword)

		emailx, _ := passwords.EncryptPassword(email, newMasterPassword)
		usernamex, _ := passwords.EncryptPassword(username, newMasterPassword)
		passwordx, _ := passwords.EncryptPassword(password, newMasterPassword)
		websitex, _ := passwords.EncryptPassword(website, newMasterPassword)
		categoryx, _ := passwords.EncryptPassword(category, newMasterPassword)
		notesx, _ := passwords.EncryptPassword(notes, newMasterPassword)

		newentries = append(newentries, helpers.Entry{
			Email:    emailx,
			Username: usernamex,
			Password: passwordx,
			Website:  websitex,
			Category: categoryx,
			Note:     notesx,
		})

		fmt.Println("Fixed password", i)

	}

	helpers.SaveEntries(newentries)

	//CHANGE THE HASH.TXT
	encryptedPassword, _ := passwords.EncryptPassword("GOPASSWORD", newMasterPassword)
	helpers.SaveTXTFile("./hash.txt", encryptedPassword)
}

func authenticateUser() (bool, string) {
	auth, pass := validateMasterPassword()
	if !auth {
		fmt.Println("INVALID MASTER PASSWORD, PLEASE RETRY")
		main()
	}
	fmt.Println("Valid master password")
	return true, pass
}

func validateMasterPassword() (bool, string) {
	filecontent := helpers.LoadTXTFile("./hash.txt")
	fmt.Println(filecontent)
	if len(filecontent) == 0 {
		//create new password
		masterPassword := helpers.GetStrInput("Create new master password:")
		encryptedPassword, _ := passwords.EncryptPassword("GOPASSWORD", masterPassword)
		helpers.SaveTXTFile("./hash.txt", encryptedPassword)
		return true, masterPassword

	} else {
		masterPassword := helpers.GetStrInput("Enter master password :")
		decryptedPassword, _ := passwords.DecryptPassword(filecontent, masterPassword)

		fmt.Println("dc", decryptedPassword)
		fmt.Println("mp", masterPassword)

		if decryptedPassword == "GOPASSWORD" {
			return true, masterPassword
		} else {
			return false, masterPassword
		}

	}

}

func main() {
	logger := helpers.Logger{Level: 1, Prefix: ""}
	authenticated, masterPassword := authenticateUser()
	if !authenticated {
		return
	}
	for true {
		logger.LogList(OPTIONS)
		option := helpers.GetIntInput()
		switch option {
		case 1:
			newEntry(&masterPassword, &logger)
		case 2:
			deleteEntry(&masterPassword)
		case 3:
			viewEntries(&masterPassword)
		case 4:
			changeMasterPassword(&masterPassword)
		case 5:
			os.Exit(0)

		}
	}
}
