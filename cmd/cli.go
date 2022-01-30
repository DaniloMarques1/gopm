package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/danilomarques1/gopm/cmd/dto"
	"github.com/danilomarques1/gopm/cmd/service"
	"golang.org/x/crypto/ssh/terminal"
)

type CLI struct {
	masterService *service.MasterService
	scanner       *bufio.Scanner
	token         string // I GUESS?
}

func NewCLI() *CLI {
	masterService := service.NewMasterService()
	scanner := bufio.NewScanner(os.Stdin)
	return &CLI{masterService: masterService, scanner: scanner}
}

func (cli *CLI) Register() {
	email := cli.getEmailFromInput()
	password, err := cli.getPasswordFromInput()
	if err != nil {
		log.Fatal(err)
	}
	registerDto := dto.MasterRegisterDto{Email: email, Pwd: password}
	if err := cli.masterService.Register(registerDto); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Master successfully created")
}

func (cli *CLI) Access() {
	email := cli.getEmailFromInput()
	password, err := cli.getPasswordFromInput()
	if err != nil {
		log.Fatal(err)
	}
	sessionDto := dto.SessionRequestDto{Email: email, Pwd: password}
	response, err := cli.masterService.Access(sessionDto)
	if err != nil {
		log.Fatal(err)
	}
	cli.token = response.Token
	cli.Shell()
}

func (cli *CLI) Shell() {
	if cli.token == "" {
		log.Fatal("You should log in first")
	}
	var input string
	for {
		fmt.Print(">> ")
		if cli.scanner.Scan() {
			input = cli.scanner.Text()
		}

		cmd, args, err := cli.parseInput(input)
		if err != nil {
			continue
		}

		switch cmd {
		case HELP:
			cli.Usage()
		case GET:
			if len(args) < 1 {
				fmt.Println("You need to provide the key of the password. See help for instructions")
				continue
			}
			// TODO do a get request for a specific password
		case SAVE:
			if len(args) < 2 {
				fmt.Println("You need to provide both the key and the password. See help for instructions.")
				continue
			}
			// TODO do a post request to save a password
		case REMOVE:
			if len(args) < 1 {
				fmt.Println("You need to provide the key of the password. See help for instructions")
				continue
			}
			// TODO do a delete request o remove a password
		case KEYS:
			// TODO do a get request for all keys
		case CLEAR:
			operatingSystem := runtime.GOOS
			var cmdToBeExecuted string

			switch operatingSystem {
			case "linux":
				cmdToBeExecuted = "clear"
			case "windows":
				cmdToBeExecuted = "cls"
			default:
				cmdToBeExecuted = "Not implemented yet"

			}

			cmd := exec.Command(cmdToBeExecuted)
			out, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(out))
		case EXIT:
			os.Exit(1)
		default:
			fmt.Println(CMD_NOT_FOUND)
			continue
		}
	}
}

// request for user input without showing whats being inputed
func (cli *CLI) getPasswordFromInput() (string, error) {
	fmt.Print("Type your password: ")
	bytes, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}
	pwd := string(bytes)
	fmt.Println() // cleaning
	return pwd, nil
}

// reads user input for an email
func (cli *CLI) getEmailFromInput() string {
	fmt.Print("Type your email address: ")
	var email string
	if cli.scanner.Scan() {
		email = cli.scanner.Text()
	}
	return email
}

// will get the user input like save password_key password_value
// and return the command, in this case save, and its arguments,
// in this case [password_key, password_value]
func (cli *CLI) parseInput(input string) (string, []string, error) {
	if len(input) == 0 {
		return "", nil, errors.New(CMD_NOT_FOUND)
	}
	splitted := strings.Split(input, " ")
	return splitted[0], splitted[1:], nil
}

// show a list of available commands
func (cli *CLI) Usage() {
	fmt.Println("The commands are used as follow")
	fmt.Println("\tget {password_name}")
	fmt.Println("\tsave {password_name} {password_itself}")
	fmt.Println("\tclear")
	fmt.Println()
	fmt.Println("List of available commands:")
	fmt.Printf("\t%v        \tShow the usage of the program\n", HELP)
	fmt.Printf("\t%v        \tHave a shell access as master\n", ACCESS)
	fmt.Printf("\t%v        \tWill retrieve a stored password. You need to provide the password name you used when you save it\n", GET)
	fmt.Printf("\t%v        \tWil save a password. You need to provide the password name and the password itself when using the command\n", SAVE)
	fmt.Printf("\t%v        \tWil remove a password. You need to provide the password name when using the command\n", REMOVE)
	fmt.Printf("\t%v        \tReturn all the passwords keys stored\n", KEYS)
	fmt.Printf("\t%v        \tExits the password manager\n", EXIT)
	fmt.Printf("\t%v        \tClears the shell\n", CLEAR)
}