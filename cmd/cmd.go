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
		break
	case REGISTER:
		// TODO
		// requisitar email
		// requisitar senha
		// enviar requisicao post para /master
		cli.Register()
		break
	default:
		cli.Usage()
	}
}
