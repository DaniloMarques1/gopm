package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func (cli *CLI) RequireMasterPassword() {
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

func (cli *CLI) getEmailFromInput() string {
	fmt.Print("Type your email address: ")
	var email string
	if cli.scanner.Scan() {
		email = cli.scanner.Text()
	}
	return email
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
