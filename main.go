package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	part2 := flag.Bool("p2", false, "-p2 to pass part2 flag to binary (may or may not support)")

	day := flag.Int("d", 1, "-d X for day X")
	flag.Parse()

	goPath, err := exec.LookPath("go")
	if err != nil {
		panic(err)
	}

	dayPath := fmt.Sprintf("./advent%d", *day)
	os.Chdir(dayPath)
	baseArgs := []string{goPath, "run", fmt.Sprintf("%s.go", dayPath)}
	if *part2 {
		baseArgs = append(baseArgs, "-p2")
	}

	cmd := exec.Cmd{
		Path:   goPath,
		Args:   baseArgs,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
