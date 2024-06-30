package tar

import (
	"os"
	"os/exec"
	"path"
)

func Extract(from, to string) error {
	os.MkdirAll(to, 0777)
	var cmd = exec.Command("tar", "-xf", from, "-C", to)
	if path.Ext(from) == ".gz" {
		cmd.Args = append(cmd.Args, "-z")
	}
	return cmd.Run()
}
