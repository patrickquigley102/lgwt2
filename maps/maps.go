package main

import (
	"errors"
)

// Dictionary stores words and their definitions.
type Dictionary map[string]string

// Search finds the definition of a word.
func (d Dictionary) Search(word string) (string, error) {
	word, found := d[word]
	if !found {
		return "", errNoWord
	}

	return word, nil
}

var (
	errNoWord          = errors.New("no word in dictionary")
	errPreExistingWord = errors.New("word already exists")
)

// Add a word to dictionary.
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if !errors.Is(err, errNoWord) {
		return errPreExistingWord
	}

	d[word] = def

	return nil
}
