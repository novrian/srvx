package main

import (
	"bufio"
	"fmt"
	// "io"
	"encoding/json"
	"github.com/novrian/srvx/pkg/server"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

const CONFIG_FILE_PATH string = "server.json"

func main() {
	listServer := ParseConfig()

	DisplayServer(listServer)
	fmt.Print("Pilihmi nomor berapa? ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Exit when input empty
	if len(input) == 1 {
		os.Exit(0)
	}

	inputInt, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		fmt.Println(err)
	}
	selected := listServer[inputInt-1]

	// @DEBUG
	fmt.Printf("%v\n", selected)

	cmd := exec.Command("ssh", selected.GenerateArgs()...)
	_, err = cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func ParseConfig() []server.Server {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	homeDir := user.HomeDir

	file, err := ioutil.ReadFile(homeDir + "/" + CONFIG_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	var servers []server.Server
	json.Unmarshal(file, &servers)

	return servers
}

func DisplayServer(listServer []server.Server) {
	headerDesc := `===============================================================================
  _______________________   ________  ___
 /   _____/\______   \   \ /   /\   \/  /
 \_____  \  |       _/\   Y   /  \     /
 /        \ |    |   \ \     /   /     \
/_______  / |____|_  /  \___/   /___/\  \
        \/         \/                 \_/
-------------------------------------------------------------------------------
SRVX - SSH Connection Manager v0.0.0
-------------------------------------------------------------------------------
`
	fmt.Printf("%s", headerDesc)

	for i, srvx := range listServer {
		fmt.Printf("%d. %s\n", i+1, srvx.Label())
	}
	fmt.Println("-------------------------------------------------------------------------------")
}
