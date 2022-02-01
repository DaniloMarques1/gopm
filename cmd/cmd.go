package cmd

// backend (request) commands
const (
	REGISTER = "register"
	GET      = "get"
	SAVE     = "save"
	REMOVE   = "remove"
	KEYS     = "keys"
	UPDATE   = "update"
)

// cli commands
const (
	ACCESS = "access"
	CLEAR  = "clear"
	EXIT   = "exit"
	HELP   = "help"
)

// errors
const (
	CMD_NOT_FOUND = "Command not found"
)

func Run(command string) {
	cli := NewCLI()
	switch command {
	case ACCESS:
		cli.Access()
	case REGISTER:
		cli.Register()
	default:
		cli.Usage()
	}
}
