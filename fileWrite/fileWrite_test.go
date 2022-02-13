package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("Momoko")
    want := "Hi, Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}