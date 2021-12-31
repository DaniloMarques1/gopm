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
	GET     = "get"
	SAVE    = "save"
	REMOVE  = "remove"
	CLEAR   = "clear"
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
	default:
		manager.help()
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
		case GET:
			if len(args) < 1 {
				fmt.Println("You need to provide the password name")
				continue
			}
			// TODO wont allow spaces
			passwordName := args[0]
			manager.getPassword(master.Id, passwordName)
		case SAVE:
			if len(args) < 2 {
				fmt.Println("You need to provide the password name along with the password itself")
				continue
			}
			pwdName, pwd := args[0], args[1]
			manager.savePassword(master.Id, pwdName, pwd)
		case REMOVE:
			if len(args) < 1 {
				fmt.Println("You need to provide the password name")
				continue
			}
			pwdName := args[0]
			manager.removePassword(master.Id, pwdName)
		case CLEAR:
			// TODO somehow clear the shell
			/*
			NOTE: this did not worked
			err := exec.Command("ls").Run()
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
			*/
		default:
			continue
		}
	}
}

func (manager *Manager) parseCmd(input string) (string, []string, error) {
	if len(input) == 0 {
		return "", nil, errors.New(CMD_NOT_FOUND)
	}

	cmdWithArgs := strings.Split(input, " ")
	return cmdWithArgs[0], cmdWithArgs[1:], nil
}

func (manager *Manager) getPassword(masterId, name string) {
	password, err := manager.passwordRepository.FindByName(masterId, name)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			fmt.Printf("The password name %v was not found\n", name)
			return
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(password.Pwd)
}

func (manager *Manager) savePassword(masterId, pwdName, pwd string) {
	password := model.Password{
		Id: uuid.NewString(),
		Name: pwdName,
		Pwd: pwd,
		MasterId: masterId,
	} 
	err := manager.passwordRepository.Save(&password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password stored successfully")
}

func (manager *Manager) removePassword(masterId, pwdName string) {
	scanner := bufio.NewScanner(os.Stdin)
	var confirmation string
	fmt.Print("Are you sure you want to delete the password? y/n ")
	if scanner.Scan() {
		confirmation = scanner.Text()
	}
	if confirmation != "y" { return }

	err := manager.passwordRepository.RemoveByName(masterId, pwdName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The password %v was removed successfully.\n", pwdName)
}

func (manager *Manager) help() {
	fmt.Println("The commands are usually used as follows")
	fmt.Println("\tget password_name")
	fmt.Println("\tsave password_name password_itself")
	fmt.Println()
	fmt.Println("List of available commands:")
	fmt.Printf("\t%v        \tShow the usage of the program\n", HELP)
	fmt.Printf("\t%v        \tHave a shell access as master\n", ACCESS)
	fmt.Printf("\t%v        \tWill retrieve a stored password. You need to provide the password name you used when you save it\n", GET)
	fmt.Printf("\t%v        \tWil save a password. You need to provide the password name and the password itself when using the command\n", SAVE)
	fmt.Printf("\t%v        \tWil remove a password. You need to provide the password name when using the command\n", REMOVE)
}
