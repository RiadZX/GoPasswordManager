package main

import (
	"GoPasswordManager/helpers"
	"GoPasswordManager/passwords"
)

var (
	OPTIONS []string = []string{
		"New Entry",
		"Delete Entry",
		"View Entries",
		"Steal Browser Passwords",
		"Exit"}
)

func main() {
	logger := helpers.Logger{Level: 1, Prefix: ""}
	for true {
		logger.LogList(OPTIONS)
		option := helpers.GetIntInput()
		switch option {
		case 1:
			entry := passwords.AddEntry()
			helpers.SaveEntry(entry)
			logger.Log("Added Password Success")
		case 2:
			passwords.ViewEntries()
			entries := passwords.DeleteEntry(helpers.LoadEntries("./passwords.json"))
			helpers.SaveEntries(entries)
		case 3:
			passwords.ViewEntries()
			logger.Log("Chosen 3")
		case 4:
			logger.Log("Chosen 4")
		}
	}
}
