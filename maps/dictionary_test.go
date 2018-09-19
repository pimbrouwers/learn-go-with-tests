package maps

import (
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "testing 123"

		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("Existing word", func(t *testing.T) {
		key := "test"
		value := "testing 123"
		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, "new testing 123")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	value := "testing 123"
	dictionary := Dictionary{key: value}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)

	if err != ErrNotFound {
		t.Errorf("expected %s (key) to be deleted", key)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "testing 123"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "testing 123"

		assertStrings(t, got, expected)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("dsdsda")

		assertError(t, err, ErrNotFound)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("Existing word", func(t *testing.T) {
		key := "test"
		value := "testing 123"
		dictionary := Dictionary{key: value}

		newValue := "new testing 123"

		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("New word", func(t *testing.T) {
		key := "test"
		value := "testing 123"
		dictionary := Dictionary{key: value}

		newKey := "test2"
		newValue := "new testing 123"

		err := dictionary.Update(newKey, newValue)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(
	t *testing.T,
	dictionary Dictionary,
	key,
	value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, value)
}

func assertError(t *testing.T, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}
