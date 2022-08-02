package test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"song.com/go_get_started/cli_cobra_example/cmd"
)

func TestAbc(t *testing.T) {
	b := bytes.NewBufferString("")

	rootCmd := cmd.RootCmd
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"login", "-u", "Song", "-p", "pwd123"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	_, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
}
