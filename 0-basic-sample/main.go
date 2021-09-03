// Package basic_sample
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-03
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		log.Fatal("unknown command\nUse run <command_name>, like `run /bin/bash` or `run echo hello`")
	}
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fatal(cmd.Run())
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
