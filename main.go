package main

import (
	"GoPasswordManager/helpers"
	"GoPasswordManager/passwords"
	"os"
)

var (
	OPTIONS []string = []string{
		"New Entry",
		"Delete Entry",
		"View Entries",
		"Steal Browser Passwords",
		"Exit",
		"Restart"}
)

func main() {
	logger := helpers.Logger{Level: 1, Prefix: ""}
	masterPassword := helpers.GetStrInput("Enter master password:")

	for true {
		logger.LogList(OPTIONS)
		option := helpers.GetIntInput()
		switch option {
		case 1:
			entry := passwords.AddEntry(masterPassword)
			helpers.SaveEntry(entry)
			logger.Log("Added Password Success")
		case 2:
			passwords.ViewEntries(masterPassword)

			entries := passwords.DeleteEntry(helpers.LoadEntries("./passwords.json"))
			if passwords.AreYouSure() {
				helpers.SaveEntries(entries)
			}

		case 3:
			entries, _ := passwords.ViewEntries(masterPassword)
			entrynum := helpers.GetIntInput()
			passwords.ViewEntry(entries[entrynum], masterPassword)

		case 4:
			logger.Log("Chosen 4")

		case 5:
			os.Exit(0)

		case 6:
			logger.Log("Restart")
		}
	}
}
