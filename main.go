package main

import (
	"GoPasswordManager/helpers"
	"GoPasswordManager/passwords"
	"fmt"
	"os"
)

var (
	OPTIONS = []string{
		"New Entry",
		"Delete Entry",
		"View Entries",
		"Change master password",
		"Export Passwords",
		"Exit",
	}

	EXPORTOPTIONS = []string{
		"CSV",
		"JSON",
	}
)

func newEntry(masterPassword *string) {
	entry := passwords.AddEntry(*masterPassword)
	helpers.SaveEntry(entry)
	helpers.Success("Added Password Success")
}

func deleteEntry(masterPassword *string) {
	entries, err := passwords.ViewEntries(*masterPassword)
	if err != nil {
		helpers.Danger(err.Error())
		return
	}
	_, err2 := passwords.DeleteEntry(entries)
	if err2 != nil {
		return
	}

	if passwords.AreYouSure() || err != nil {
		helpers.SaveEntries(entries)
	} else {
		helpers.Danger("error found deletenetry")
	}
}

func viewEntries(masterPassword *string) {
	entries, err := passwords.ViewEntries(*masterPassword)
	if err != nil {
		helpers.Danger(err.Error())
		return
	}
	if len(entries) == 0 {
		helpers.Danger(err.Error())
		return
	}
	entrynum := helpers.GetIntInput("Enter entry index to view")
	if entrynum > len(entries)-1 || entrynum < 0 {
		helpers.Danger("Invalid Input, Entry does not exist")
		return
	}
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

	*masterPassword = newMasterPassword
}

func authenticateUser() (bool, string) {
	auth, pass := validateMasterPassword()
	if !auth {
		helpers.Danger("INVALID MASTER PASSWORD, PLEASE RETRY")
		main()
	}

	return true, pass
}

func validateMasterPassword() (bool, string) {
	filecontent := helpers.LoadTXTFile("./hash.txt")

	if len(filecontent) == 0 {
		//create new password
		masterPassword := helpers.GetStrInput("Create new master password")
		encryptedPassword, _ := passwords.EncryptPassword("GOPASSWORD", masterPassword)
		helpers.SaveTXTFile("./hash.txt", encryptedPassword)
		return true, masterPassword

	} else {
		masterPassword := helpers.GetStrInput("Enter master password")
		decryptedPassword, _ := passwords.DecryptPassword(filecontent, masterPassword)

		if decryptedPassword == "GOPASSWORD" {
			return true, masterPassword
		} else {
			return false, masterPassword
		}

	}

}

func exportEntries() {

	helpers.LogList(EXPORTOPTIONS)
	option := helpers.GetIntInput("Choose an option")
	switch option {
	case 1:
		passwords.ExportCSV()
		helpers.Success("Export to CSV Success")
	case 2:
		passwords.ExportJSON()
		helpers.Success("Export to JSON Success")

	default:
		helpers.Danger("Invalid Input, back to main menu")
	}
}

func printMenu() {
	logo := helpers.LoadTXTFile("./data/menu.txt")
	helpers.LogoPrint(logo)
}

func main() {

	printMenu()
	authenticated, masterPassword := authenticateUser()
	if !authenticated {
		return
	}
	for true {
		helpers.LogList(OPTIONS)
		option := helpers.GetIntInput("Choose an option")
		switch option {
		case 1:
			newEntry(&masterPassword)
		case 2:
			deleteEntry(&masterPassword)
		case 3:
			viewEntries(&masterPassword)
		case 4:
			changeMasterPassword(&masterPassword)
		case 5:
			exportEntries()
		case 6:
			os.Exit(0)
		default:
			helpers.Danger("Invalid Input")
		}
	}
}
