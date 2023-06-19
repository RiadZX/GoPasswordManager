package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadJson(filepath string) []Entry {
	jsonFile, err := os.Open(filepath)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var entries []Entry
	json.Unmarshal(byteValue, &entries)

	return entries
}

func saveEntries(entries []Entry) {
	file, _ := json.MarshalIndent(entries, "", " ")

	_ = os.WriteFile("./passwords.json", file, 0644)
}

func SavePassword(entry Entry) {
	logger := Logger{Level: 1, Prefix: ""}
	logger.Log("Saved Password")
	var currentEntries []Entry = loadJson("./passwords.json")
	entries := append(currentEntries, entry)
	saveEntries(entries)
}
