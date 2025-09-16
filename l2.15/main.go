package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

func main() {
	signal.Ignore(os.Interrupt)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nExit")
			os.Exit(0)
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if strings.Contains(input, "|") {
			executePipeline(input)
			continue
		}

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		if isBuiltinCommand(parts[0]) {
			executeBuiltin(parts)
			continue
		}

		executeExternal(parts)
	}
}

func isBuiltinCommand(cmd string) bool {
	builtins := []string{"cd", "pwd", "echo", "kill", "ps"}
	for _, b := range builtins {
		if cmd == b {
			return true
		}
	}
	return false
}

func executeBuiltin(parts []string) {
	switch parts[0] {
	case "cd":
		if len(parts) < 2 {
			fmt.Fprintln(os.Stderr, "cd: missing argument")
			return
		}
		err := os.Chdir(parts[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "cd:", err)
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "pwd:", err)
			return
		}
		fmt.Println(dir)
	case "echo":
		fmt.Println(strings.Join(parts[1:], " "))
	case "kill":
		if len(parts) < 2 {
			fmt.Fprintln(os.Stderr, "kill: missing PID")
			return
		}
		cmd := exec.Command("kill", parts[1])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func executeExternal(parts []string) {
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func executePipeline(input string) {
	commands := strings.Split(input, "|")
	var cmds []*exec.Cmd

	for _, cmdStr := range commands {
		parts := strings.Fields(strings.TrimSpace(cmdStr))
		if len(parts) == 0 {
			continue
		}
		cmds = append(cmds, exec.Command(parts[0], parts[1:]...))
	}

	for i := 0; i < len(cmds)-1; i++ {
		stdout, err := cmds[i].StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Pipe error:", err)
			return
		}
		cmds[i+1].Stdin = stdout
	}

	cmds[len(cmds)-1].Stdout = os.Stdout
	cmds[len(cmds)-1].Stderr = os.Stderr

	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Start error:", err)
			return
		}
	}

	for _, cmd := range cmds {
		cmd.Wait()
	}
}
