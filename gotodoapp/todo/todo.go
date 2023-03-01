package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Item struct {
	Text string
}

func SavedItems(filename string, items []Item) error {
	//Converting struct data to json data.
	jsonData, err := json.Marshal(items)
	if err != nil {
		return err
	}
	//Print data on console.
	//fmt.Println(string(jsonData))
	//The os.Create creates or truncates the named file.
	//If the file already exists, it is truncated.
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	// The File.Write writes n bytes to a file.
	// The File.WriteAt writes n bytes to a file starting at
	// the specified byte offset.
	_, err2 := f.WriteString(string(jsonData))
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("done! Data inserted !")
	defer f.Close()
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	content, err := os.ReadFile(filename)
	if err != nil {

		//fmt.Println("Error in file opening", err) <- this message for debug purpose
		return []Item{}, nil
	}
	var items []Item
	//converting json data to Item struct formate.
	if err := json.Unmarshal(content, &items); err != nil {
		//fmt.Println("Error while parsing Data.")
		return []Item{}, err
	}
	return items, nil
}
