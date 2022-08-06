package cmd

import (
	"bytes"
	"testing"

	"github.com/mattn/go-shellwords"
)

func TestSelectClip(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "clip-list select hoge", want: "Arguments are not required"},
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
