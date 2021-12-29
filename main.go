package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/danilomarques1/gopm/model"
	"github.com/google/uuid"
	"github.com/danilomarques1/gopm/repository"

	_ "github.com/mattn/go-sqlite3"
)

const TABLES = `
	CREATE TABLE IF NOT EXISTS master(
		id VARCHAR(36) PRIMARY KEY,
		master_pwd VARCHAR(100) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS password(
		id VARCHAR(36) PRIMARY KEY,
		master_id VARCHAR(36) NOT NULL,
		name VARCHAR(50) NOT NULL UNIQUE,
		pwd VARCHAR(100) NOT NULL,
		FOREIGN KEY(master_id) REFERENCES master(id)
	);
`

// commands
const (
	HELP    = "help"
	ACCESS  = "access"
)

// errors
const (
	CMD_NOT_FOUND = "Command not found"
)

type Manager struct {
	masterRepository   model.MasterRepository
	passwordRepository model.PasswordRepository
}

func NewManager(masterRepository model.MasterRepository,
	passwordRepository model.PasswordRepository) *Manager {
	return &Manager{
		masterRepository: masterRepository,
		passwordRepository: passwordRepository,
	}
}

func main() {
	db, err := sql.Open("sqlite3", "gopm.db")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(TABLES); err != nil {
		log.Fatal(err)
	}

	masterRepository := repository.NewMasterRepository(db)
	passwordRepository := repository.NewPasswordRepository(db)
	manager := NewManager(masterRepository, passwordRepository)

	if len(os.Args) == 1 {
		log.Printf("You need to give the command. See help for details")
		os.Exit(1)
	}

	manager.Run()
}

func (manager *Manager) Run() {
	command, _ := manager.parseCmdArgs()
	switch command {
	case HELP:
		manager.help()
	case ACCESS:
		manager.requireMasterPwd()
	}
}

func (manager *Manager) parseCmdArgs() (string, []string) {
	args := os.Args[1:]
	return args[0], args[1:]
}

func (manager *Manager) requireMasterPwd() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type the master password>> ")
	var pwd string
	if scanner.Scan() {
		pwd = scanner.Text()
	}
	master, err := manager.masterRepository.FindByPassword(pwd)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			fmt.Print("No master found associated with the given password. Would you like to create a master with this password? y/n ")
			var answer string
			if scanner.Scan() {
				answer = scanner.Text()
			}

			if answer == "y" {
				if err := manager.createMaster(pwd); err != nil {
					log.Fatal(err)
				}
				fmt.Println("Master successfully created")
			}
			return
		} else {
			log.Fatal(err)
		}
	}

	manager.Shell(master)
}

func (manager *Manager) createMaster(pwd string) error {
	master := model.Master{Id: uuid.NewString(), MasterPwd: pwd}
	err := manager.masterRepository.Save(&master)
	return err
}

func (manager *Manager) Shell(master *model.Master) {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		fmt.Print(">> ")
		if scanner.Scan() {
			input = scanner.Text()
		}

		cmd, args, err := manager.parseCmd(input)
		if err != nil {
			log.Fatal(err)
		}
		switch cmd {
		case HELP:
			manager.help()
		default:
			continue
		}

		fmt.Println(cmd)
		fmt.Println(args)
		joined := strings.Join(args, " ")
		fmt.Println(joined)
	}
}

func (manager *Manager) parseCmd(input string) (string, []string, error) {
	if len(input) == 0 {
		return "", nil, errors.New(CMD_NOT_FOUND)
	}

	cmdWithArgs := strings.Split(input, " ")
	return cmdWithArgs[0], cmdWithArgs[1:], nil
}

func (manager *Manager) help() {
	fmt.Println("List of available commands:")
	fmt.Println("\thelp   - \tShow the usage of the program")
	fmt.Println("\taccess - \tHave a shell acess as master")
}
