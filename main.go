package main

import (
	"errors"
	"log"
	"os"

	"github.com/danilomarques1/gopm/cmd"
)

func main() {
	command, err := parseCmdArgs()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Run(command)
	/*
	// manager will have pointers to services 
	manager := NewManager(masterRepository, passwordRepository)

	if len(os.Args) == 1 {
		log.Fatal("You need to give a command. See help for details")
	}

	manager.Run()
	*/
}

func parseCmdArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("You should provide a command. Type help for instructions")
	}
	command := os.Args[1]
	return command, nil
}

// will return the database fileName that wil be
// the instalation location + the fileName itself
//func getFileName() (string, error) {
//	path, err := os.Executable()
//	if err != nil {
//		return "", err
//	}
//	return path + ".db", nil
//}
//
//func (manager *Manager) Run() {
//	command, _ := manager.parseCmdArgs()
//	switch command {
//	case ACCESS:
//		manager.requireMasterPwd()
//	default:
//		manager.help()
//	}
//}
//
//
//func (manager *Manager) requireMasterPwd() {
//	fmt.Print("Type the master password>> ")
//	pwdBytes, err := terminal.ReadPassword(syscall.Stdin)
//	if err != nil {
//		log.Fatal(err)
//	}
//	pwd := string(pwdBytes)
//	fmt.Println() // clears the buffer
//
//	master, err := manager.masterRepository.FindByPassword(pwd)
//	if err != nil {
//		if errors.Is(sql.ErrNoRows, err) {
//			fmt.Print("No master found associated with the given password. Would you like to create a master with this password (y/n)? ")
//			var answer string
//			scanner := bufio.NewScanner(os.Stdin)
//			if scanner.Scan() {
//				answer = scanner.Text()
//			}
//
//			if answer == "y" {
//				if err := manager.createMaster(pwd); err != nil {
//					log.Fatal(err)
//				}
//				fmt.Println("Master successfully created")
//			}
//			return
//		} else {
//			log.Fatal(err)
//		}
//	}
//
//	manager.Shell(master)
//}
//
//func (manager *Manager) createMaster(pwd string) error {
//	master := model.Master{Id: uuid.NewString(), MasterPwd: pwd}
//	err := manager.masterRepository.Save(&master)
//	return err
//}
//
//func (manager *Manager) Shell(master *model.Master) {
//	scanner := bufio.NewScanner(os.Stdin)
//	var input string
//	for {
//		fmt.Print(">> ")
//		if scanner.Scan() {
//			input = scanner.Text()
//		}
//
//		cmd, args, err := manager.parseCmd(input)
//		if err != nil {
//			continue
//		}
//
//		switch cmd {
//		case HELP:
//			manager.help()
//		case GET:
//			if len(args) < 1 {
//				fmt.Println("You need to provide the password name")
//				continue
//			}
//			// TODO wont allow spaces
//			passwordName := args[0]
//			manager.getPassword(master.Id, passwordName)
//		case SAVE:
//			if len(args) < 2 {
//				fmt.Println("You need to provide the password name along with the password itself")
//				continue
//			}
//			pwdName, pwd := args[0], args[1]
//			manager.savePassword(master.Id, pwdName, pwd)
//		case REMOVE:
//			if len(args) < 1 {
//				fmt.Println("You need to provide the password name")
//				continue
//			}
//			pwdName := args[0]
//			manager.removePassword(master.Id, pwdName)
//		case KEYS:
//			manager.showKeys(master.Id)
//		case CLEAR:
//			operatingSystem := runtime.GOOS
//			var cmdToBeExecuted string
//
//			switch operatingSystem {
//			case "linux":
//				cmdToBeExecuted = "clear"
//			case "windows":
//				cmdToBeExecuted = "cls"
//			default:
//				cmdToBeExecuted = "Not implemented yet"
//
//			}
//
//			cmd := exec.Command(cmdToBeExecuted)
//			out, err := cmd.Output()
//			if err != nil {
//				log.Fatal(err)
//			}
//			fmt.Print(string(out))
//		case EXIT:
//			os.Exit(1)
//		default:
//			fmt.Println(CMD_NOT_FOUND)
//			continue
//		}
//	}
//}
//
//func (manager *Manager) parseCmd(input string) (string, []string, error) {
//	if len(input) == 0 {
//		return "", nil, errors.New(CMD_NOT_FOUND)
//	}
//
//	cmdWithArgs := strings.Split(input, " ")
//	return cmdWithArgs[0], cmdWithArgs[1:], nil
//}
//
//func (manager *Manager) getPassword(masterId, name string) {
//	password, err := manager.passwordRepository.FindByName(masterId, name)
//	if err != nil {
//		if errors.Is(sql.ErrNoRows, err) {
//			fmt.Printf("The password name %v was not found\n", name)
//			return
//		} else {
//			log.Fatal(err)
//		}
//	}
//	fmt.Println(password.Pwd)
//}
//
//func (manager *Manager) savePassword(masterId, pwdName, pwd string) {
//	password := model.Password{
//		Id:       uuid.NewString(),
//		Name:     pwdName,
//		Pwd:      pwd,
//		MasterId: masterId,
//	}
//	err := manager.passwordRepository.Save(&password)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Password stored successfully")
//}
//
//func (manager *Manager) removePassword(masterId, pwdName string) {
//	scanner := bufio.NewScanner(os.Stdin)
//	var confirmation string
//	fmt.Print("Are you sure you want to delete the password (y/n)? ")
//	if scanner.Scan() {
//		confirmation = scanner.Text()
//	}
//	if confirmation != "y" {
//		return
//	}
//
//	err := manager.passwordRepository.RemoveByName(masterId, pwdName)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("The password %v was removed successfully.\n", pwdName)
//}
//
//func (manager *Manager) showKeys(masterId string) {
//	passwords, err := manager.passwordRepository.FindAll(masterId)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, password := range passwords {
//		fmt.Println(password.Name)
//	}
//}
//
//func (manager *Manager) help() {
//	fmt.Println("The commands are usually used as follows")
//	fmt.Println("\tget password_name")
//	fmt.Println("\tsave password_name password_itself")
//	fmt.Println("\tclear")
//	fmt.Println()
//	fmt.Println("List of available commands:")
//	fmt.Printf("\t%v        \tShow the usage of the program\n", HELP)
//	fmt.Printf("\t%v        \tHave a shell access as master\n", ACCESS)
//	fmt.Printf("\t%v        \tWill retrieve a stored password. You need to provide the password name you used when you save it\n", GET)
//	fmt.Printf("\t%v        \tWil save a password. You need to provide the password name and the password itself when using the command\n", SAVE)
//	fmt.Printf("\t%v        \tWil remove a password. You need to provide the password name when using the command\n", REMOVE)
//	fmt.Printf("\t%v        \tReturn all the passwords keys stored\n", KEYS)
//	fmt.Printf("\t%v        \tExits the password manager\n", EXIT)
//	fmt.Printf("\t%v        \tClears the shell\n", CLEAR)
//}
