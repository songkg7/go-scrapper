package dict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errWordExists = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		return errWordExists
	}
	if err == errNotFound {
		d[word] = def
	}
	return nil
}
