package main

import (
	"errors"
)

const (
	ErrorWordNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrorWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrorWordNotFound):
		d[word] = definition
	case err == nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrorWordNotFound):
		return ErrorWordDoesNotExist
	case err == nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrorWordNotFound):
		return ErrorWordDoesNotExist
	case err == nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

func main() {

}
