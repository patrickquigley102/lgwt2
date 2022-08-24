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

// Add a word to dictionary.
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if !errors.Is(err, errNoWord) {
		return errPreExistingWord
	}

	d[word] = def

	return nil
}

// Update the definition of a word in dictionary.
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	if errors.Is(err, errNoWord) {
		return err
	}

	d[word] = def

	return nil
}

// Delete a word from dictionary.
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if errors.Is(err, errNoWord) {
		return err
	}

	delete(d, word)

	return nil
}

type dictionaryErr string

const (
	errNoWord          = dictionaryErr("no word in dictionary")
	errPreExistingWord = dictionaryErr("word already exists")
)

func (e dictionaryErr) Error() string {
	return string(e)
}
