package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")

		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		want := ErrNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New add", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("test", "new word")

		want := "new word"

		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("Existing add", func(t *testing.T) {
		word := "word"
		definition := "means something"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "collision")

		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "word"
		definition := "means something"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		got, _ := dictionary.Search(word)

		assertStrings(t, got, newDefinition)
	})

	t.Run("Missing word", func(t *testing.T) {
		word := "word"
		definition := "means something"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update("missing", newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
