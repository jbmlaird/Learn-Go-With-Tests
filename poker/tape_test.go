package poker

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	t.Run("overwrite file contents", func(t *testing.T) {
		file, clean := createTempFile(t, "12345")
		defer clean()

		tape := &tape{*file}

		tape.Write([]byte("abc"))

		file.Seek(0, 0)
		newFileContents, _ := ioutil.ReadAll(file)

		got := string(newFileContents)
		want := "abc"

		if got != want {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
}
