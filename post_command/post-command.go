package postCommand

import (
	"os"
	"strings"
)

// PostCommand struct
type postCommand struct {
	commands []string
}

var p = postCommand{}

// ChangeDir sets the directory to cd to after command exits
func ChangeDir(dirPath string) {
	p.commands = append(p.commands, "cd "+dirPath)
}

// RunCommand adds a command to the list of commands that will run after it program exists
func RunCommand(command string) {
	p.commands = append(p.commands, command)
}

// Write writes to the file descriptor
func Write() {
	fd := os.NewFile(8, "fd")
	postCommandString := writeString() + "\n"
	fd.Write([]byte(postCommandString))
	fd.Close()
}

func writeString() string {
	return strings.Join(p.commands, "\n")
}
