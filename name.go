package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	firstname, lastname, middlename, firstInitial, lastInitial, middleInitial string
	aliases                                                                   []string
}

func (n *Name) Init(s map[string]string) {

	if s["firstname"] != "" {
		n.firstname = s["firstname"]
		n.firstInitial = string(s["firstname"][0])
	}
	if s["lastname"] != "" {
		n.lastname = s["lastname"]
		n.lastInitial = string(s["lastname"][0])
	}
	if s["middlename"] != "" {
		n.middlename = s["middlename"]
		n.middleInitial = string(s["middlename"][0])
	}

}

func getPeople(file *os.File, morphedName chan map[string]string) {
	scanning := bufio.NewScanner(file)
	defer close(morphedName)
	for scanning.Scan() {
		scanned_name := strings.Fields(scanning.Text())
		mapped_name := make(map[string]string)
		switch len(scanned_name) {
		case 1:
			mapped_name["firstname"] = scanned_name[0]
		case 2:
			mapped_name["firstname"] = scanned_name[0]
			mapped_name["lastname"] = scanned_name[1]
		case 3:
			mapped_name["firstname"] = scanned_name[0]
			mapped_name["lastname"] = scanned_name[2]
			mapped_name["middlename"] = scanned_name[1]
		default:
			fmt.Fprintf(os.Stderr, "Skipped invalid line: %v\n", scanned_name)
			continue
		}
		morphedName <- mapped_name
	}
	if err := scanning.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Scanning error: %v\n", err)
	}
}
