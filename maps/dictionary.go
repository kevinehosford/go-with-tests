package maps

// Dictionary is a thin wrapper around map
type Dictionary map[string]string

// DictionaryErr is a wrapper around string that implements the error interface
type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

const (
	// ErrNotFound is not found
	ErrNotFound = DictionaryErr("not found")
	// ErrWordExists is exists
	ErrWordExists = DictionaryErr("word exists")
	// ErrWordDoesNotExist blah
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

// Search searches for a type
func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

// Add adds a new word
func (d Dictionary) Add(key string, value string) error {
	d[key] = value

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update updates a word
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}
