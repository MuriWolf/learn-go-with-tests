package main

const (
	NotFound       = DictionaryError("could not find word in dictionary")
	AlreadyDefined = DictionaryError("word is already defined")
	NotDefined     = DictionaryError("word is not defined, operation cannot continue")
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

func (d Dictionary) Update(word, definition string) error {
	_, error := d.Search(word)

	switch error {
	case NotFound:
		return NotDefined
	case nil:
		d[word] = definition
	default:
		return error
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, error := d.Search(word)

	switch error {
	case NotFound:
		return NotDefined
	case nil:
		delete(d, word)
	default:
		return error
	}

	return nil
}
