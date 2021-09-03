// Package basic_sample
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-03
//go:build linux && !windows
// +build linux,!windows

package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

const (
	osRootFS = "../os-root-fs/ubuntu"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "container":
		container()
	default:
		log.Fatal("unknown command\nUse run <command_name>, like `run /bin/bash` or `run echo hello`")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"container"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Cloneflags is only available in Linux
	// CLONE_NEWUTS namespace isolates uts (hostname)
	// CLONE_NEWPID namespace isolates processes
	// CLONE_NEWNS namespace isolates mounts
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUT | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	fatal(cmd.Run())
}

func container() {
	fmt.Printf("Container -> PID: %d :: running command: [%v]\n", os.Getpid(), os.Args[2:])

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot(osRootFS))
	// Remember to extract the ubuntu-fs.tar.gz to the ./ubuntu directory
	must(os.Chdir(osRootFS))
	// Mount /proc inside container so that `ps` command works
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	// Mount a temporary filesystem
	must(syscall.Mount("thing", "new_temp", "tmpfs", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fatal(cmd.Run())

	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("thing", 0))
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
