package cmd

import (
	"bytes"
	"testing"

	"github.com/mattn/go-shellwords"
)

func TestAddClip(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "clip-list add hoge", want: "Added hoge.\n"},
		{command: "clip-list add", want: "Requires one argument"},
		{command: "clip-list add \"This is test\"", want: "Added This is test.\n"},
		{command: "clip-list add hoge fuga", want: "Requires one argument"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot("../.clipList")
		cmd.SetOutput(buf)
		cmdArgs, err := shellwords.Parse(c.command)

		if err != nil {
			t.Fatalf("args parse error: %v\n", err)
		}
		cmd.SetArgs(cmdArgs[1:])
		err = cmd.Execute()

		if err != nil {
			if c.want != err.Error() {
				t.Errorf("expected: %v, actual: %s", c.want, err.Error())
			}
		} else {
			output := buf.String()
			if c.want != output {
				t.Errorf("expected: %v, actual: %s", c.want, output)
			}
		}
	}

}
