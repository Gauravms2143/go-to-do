package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	//"github.com/hashicorp/hcl/hcl/strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
}

func (i *Item) Label() string {
	//convert Integer to ASCII.
	return strconv.Itoa(i.position) + "."
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 2 {
		return "(2)"
	}
	return " "
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
	// Note : We are taking &items instead of items and storing accordingly.
	if err := json.Unmarshal(content, &items); err != nil {
		//fmt.Println("Error while parsing Data.")
		return []Item{}, err
	}

	for index, _ := range items {
		items[index].position = index + 1
	}

	return items, nil
}

//ByPri implements the sort.Interface for []Item based
//on priority and position feild.
//This is a code snippet that defines a custom sort order for a slice
// of Item structs based on their priority and position fields.

// Below New type ByPri, which is a slice of Item structs.
type ByPri []Item

// //Len() defines a method on the ByPri type,
// which returns the length of the slice.
func (s ByPri) Len() int { return len(s) }

// Swap() defines a method on the ByPri type, which
// swaps the elements at the given indices i and j in the slice.
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less() defines a method on the ByPri type,
// which compares two elements in the slice at indices i and j,
// and returns a boolean indicating whether the element at
// index i is less than the element at index j.
func (s ByPri) Less(i, j int) bool {
	//first compare Priority
	if s[i].Priority == s[j].Priority {
		//if priority same return result based on position.
		return s[i].position < s[j].position
	}
	return s[i].Priority > s[j].Priority
}
