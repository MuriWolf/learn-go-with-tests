package main

import "errors"

var ErrWordDefinitionNotFound = errors.New("could not find word in dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordDefinitionNotFound
	}
	return definition, nil
}
