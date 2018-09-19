package maps

// ErrNotFound represents an failed Dictionary key lookup
const (
	ErrNotFound         = DictionaryErr("could not find word")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
	ErrWordExists       = DictionaryErr("word exists")
)

// DictionaryErr represents all errors related to the Dictionary type
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary represents a collection of string key/value pairs
type Dictionary map[string]string

// Add appends the item to the underlying dictionary
func (d Dictionary) Add(key, value string) error {
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

// Delete removes item from dictionary
func (d Dictionary) Delete(key string) {
	delete(d, key)
}

// Search returns matching value from map for provided key
func (d Dictionary) Search(key string) (string, error) {
	match, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return match, nil
}

// Update sets the value at the given key
func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = newValue
	default:
		return err
	}

	return nil
}
