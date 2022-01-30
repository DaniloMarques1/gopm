package cmd

// backend (request) commands
const (
	REGISTER = "register"
	GET      = "get"
	SAVE     = "save"
	REMOVE   = "remove"
	KEYS     = "keys"
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
		// TODO
		cli.Access()
		break
	case REGISTER:
		cli.Register()
		break
	default:
		cli.Usage()
	}
}
