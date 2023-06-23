package passwords

import (
	"GoPasswordManager/helpers"
	"errors"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func AddEntry(masterPassword string) helpers.Entry {
	println("STARTING ADDING")
	email, _ := EncryptPassword(helpers.GetStrInput("Email"), masterPassword)
	username, _ := EncryptPassword(helpers.GetStrInput("Username"), masterPassword)
	password, _ := EncryptPassword(helpers.GetPasswordInput(), masterPassword)
	website, _ := EncryptPassword(helpers.GetStrInput("Website"), masterPassword)
	category, _ := EncryptPassword(helpers.GetStrInput("Category"), masterPassword)
	notes, _ := EncryptPassword(helpers.GetStrInput("Note"), masterPassword)

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
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Entries")
	t.AppendHeader(table.Row{"#", "Email", "Username", "Password", "Website", "Category", "Note"})

	if len(entries) == 0 {
		entries := []helpers.Entry{helpers.Entry{}}
		return entries, errors.New("no entries found")
	}
	for i, v := range entries {
		email, _ := DecryptPassword(v.Email, masterPassword)
		username, _ := DecryptPassword(v.Username, masterPassword)
		password, _ := DecryptPassword(v.Password, masterPassword)
		website, _ := DecryptPassword(v.Website, masterPassword)
		category, _ := DecryptPassword(v.Category, masterPassword)
		notes, _ := DecryptPassword(v.Note, masterPassword)
		t.AppendRow(table.Row{i, email, username, password, website, category, notes})
	}

	t.AppendSeparator()

	t.AppendFooter(table.Row{"", "", "Total Entries", len(entries)})
	t.SetStyle(table.StyleRounded)
	t.Render()
	return entries, nil
}

func ViewEntry(entry helpers.Entry, masterPassword string) {
	email, _ := DecryptPassword(entry.Email, masterPassword)
	username, _ := DecryptPassword(entry.Username, masterPassword)
	password, _ := DecryptPassword(entry.Password, masterPassword)
	website, _ := DecryptPassword(entry.Website, masterPassword)
	category, _ := DecryptPassword(entry.Category, masterPassword)
	notes, _ := DecryptPassword(entry.Note, masterPassword)

	helpers.LogEntry(email, username, password, website, category, notes)
}

func DeleteEntry(entries []helpers.Entry) ([]helpers.Entry, error) {
	//Remove element at index from entries
	index := helpers.GetIntInput("Enter entry index to delete")
	if index > len(entries)-1 || index < 0 {
		helpers.Danger("Invalid Input, Entry does not exist")
		return []helpers.Entry{}, errors.New("Invalid Input")
	}
	return append(entries[:index], entries[index+1:]...), nil
}

func ExportCSV() {

}

func ExportJSON() {

}

func AreYouSure() bool {
	//ask the user for confirmation, write the code for me
	resp := helpers.GetStrInput("Are you sure? (y/n): ")
	if resp == "y" || resp == "Y" || resp == "yes" || resp == "YES" || resp == "1" {
		return true
	}

	return false
}
