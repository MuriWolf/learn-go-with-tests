package main

const (
	NotFound       = DictionaryError("could not find word in dictionary")
	AlreadyDefined = DictionaryError("word is already defined")
)

type DictionaryError string

func (de DictionaryError) Error() string {
	return string(de)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", NotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, error := d.Search(word)

	switch error {
	case NotFound:
		d[word] = definition
	case nil:
		return AlreadyDefined
	default:
		return error
	}

	return nil
}
