package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/danilomarques1/gopm/model"
	"github.com/danilomarques1/gopm/repository"

	_ "github.com/mattn/go-sqlite3"
)

const tables = `
	CREATE TABLE IF NOT EXISTS master(
		id PRIMARY KEY INTEGER AUTOINCREMENT
		master_pwd VARCHAR(100) NOT NULL,
	);
	CREATE TABLE IF NOT EXISTS password(
		id PRIMARY KEY INTEGER AUTOINCREMENT
		master_id INTEGER,
		name VARCHAR(50) NOT NULL,
		pwd VARCHAR(100) NOT NULL,
		FOREIGN KEY(master_id) REFERENCES master(id)
	);
`

// commands
const (
	HELP    = "help"
	ACCESS  = "access"
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

	masterRepository := repository.NewMasterRepository(db)
	passwordRepository := repository.NewPasswordRepository(db)
	manager := NewManager(masterRepository, passwordRepository)

	if len(os.Args) == 1 {
		log.Printf("You need to give the command. See help for details")
		os.Exit(1)
	}

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
}

func (manager *Manager) help() {
	fmt.Println("List of available commands:")
	fmt.Println("\thelp   - \tShow the usage of the program")
	fmt.Println("\taccess - \tHave a shell acess as master")
}
