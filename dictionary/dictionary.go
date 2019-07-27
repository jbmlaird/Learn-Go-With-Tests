package dictionary

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrWordNotExisting = DictionaryErr("could not find word in dictionary")
	ErrWordExists      = DictionaryErr("word already exists in dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotExisting
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// You need to check for 3 cases:
	// word exists
	// word doesn't exist
	// added word
	// The below only checks for word exists or added word
	_, err := d.Search(word)

	switch err {
	case ErrWordNotExisting:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotExisting:
		return err
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
