package dictionary

import "testing"

func TestSearch(t *testing.T) {
	const knownKey = "test"
	const definition = "this is just a test"
	dictionary := Dictionary{knownKey: definition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(knownKey)
		want := definition

		assertStrings(t, got, want, knownKey)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("wtf bro")

		assertError(t, got, ErrWordNotExisting)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "david"
		definition := "godwin"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word to dictionary", func(t *testing.T) {
		word := "test"
		definition := "definition"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Updated value of word that exists", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Update value that doesn't exist. Throw error", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordNotExisting)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word that exists", func(t *testing.T) {
		word := "word"
		dictionary := Dictionary{word: "definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrWordNotExisting {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertStrings(t *testing.T, got, want string, given string) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q but got %q, given %q", want, got, given)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Fatal("Expected error not thrown")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("Got %q wanted %q", got, definition)
	}
}
