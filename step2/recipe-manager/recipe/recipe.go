package recipe

import (
	"fmt"
)

var db = make(map[string]string)

func Add(name string) {
	db[name] = name
	fmt.Printf("INFO - '%s' added", name)
}

func Remove(name string) {
	if _, value := db[name]; value {
		delete(db, name)
	} else {
		fmt.Printf("ERROR - '%s' not found in db!", name)
	}
}

func Count() int {
	return len(db)
}
