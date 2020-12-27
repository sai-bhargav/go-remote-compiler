package dockersandbox

import (
	"fmt"
	"log"
	"os/exec"
)

type Sandbox struct {
	Lang    string
	Code    string
	Input   string
	Command string
}

type Response struct {
	output string
	status string
}

func Run(language string, code string, inputs string) string {

	sandbox := prepare(language, code, inputs)

	out := execute(sandbox)
	cleanup()

	return out
}

func prepare(language string, code string, inputs string) Sandbox {

	command := fmt.Sprintf("%s %s %s", language, code, inputs)
	return Sandbox{language, code, inputs, command}
}

func execute(sandbox Sandbox) string {

	cmdStr := "docker run --rm -v $(pwd)/payloads:/app remote-code-compiler " + sandbox.Command

	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()

	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}

	return string(out)
}

func cleanup() {
	// to be implemented
}
