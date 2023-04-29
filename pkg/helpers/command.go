package helpers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Verbose bool
	Folder  string
}

func NewCommand(verbose bool, folder string) *Command {
	return &Command{Folder: folder, Verbose: verbose}
}

func (c *Command) Run(command string) {
	os.Chdir(c.Folder)

	fmt.Printf("       %s  %s\n", Green("run"), command)
	parts := strings.Split(command, " ")

	cmd := exec.Command(parts[0], parts[1:]...)

	if !c.Verbose {
		cmd.Run()

	} else {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Printf("    %s  %s\n", Red("output"), err)
			return
		}

		err = cmd.Start()
		if err != nil {
			fmt.Printf("    %s  %s\n", Red("output"), err)
			return
		}

		scanner := bufio.NewScanner(stdout)

		for scanner.Scan() {
			fmt.Printf("    %s  %s\n", Yellow("output"), scanner.Text())
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Printf("    %s  %s\n", Red("output"), err)
			return
		}
	}
}
