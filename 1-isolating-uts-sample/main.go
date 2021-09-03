// Package basic_sample
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-03
//go:build linux && !windows
// +build linux,!windows

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
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
	fmt.Printf("Running %v as pid %d \n", os.Args[2:], os.Getegid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Cloneflags is only available in Linux
	// CLONE_NEWUTS namespace isolates uts (hostname)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUT,
	}

	fatal(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
