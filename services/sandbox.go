package dockersandbox

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type Sandbox struct {
	Language   string `json:"lang"`
	Code       string `json:"code"`
	TestCases  string `json:"tc"`
	Identifier string `json:"identifier"`
	Command    string
}

func Run(sandbox Sandbox) string {

	prepare(&sandbox)
	out := execute(sandbox)
	cleanup()

	return out
}

func prepare(sandbox *Sandbox) {
	createVolume(sandbox)

	languages := map[string]map[string]string{}

	file, _ := ioutil.ReadFile("config/languages.yml")

	err := yaml.Unmarshal(file, languages)
	if err != nil {
		log.Fatal(err)
	}

	if len(languages[sandbox.Language]) == 0 {
		log.Fatal("Language not configured")
	}

	sandbox.Command = languages[sandbox.Language]["command"]
}

func execute(sandbox Sandbox) string {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	sourceVolume := fmt.Sprintf("%s/payloads/%s", pwd, sandbox.Identifier)
	dockerCommand := fmt.Sprintf("docker run --rm -v %s:/app remote-code-compiler %s", sourceVolume, sandbox.Command)

	out, err := exec.Command("/bin/sh", "-c", dockerCommand).Output()

	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}

	return string(out)
}

func createVolume(sandbox *Sandbox) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	requestVolume := fmt.Sprintf("%s/payloads/%s", pwd, sandbox.Identifier)

	err = os.MkdirAll(requestVolume, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// writing the submitted code into a file with name code.language
	codeFilePath := fmt.Sprintf("%s/code.%s", requestVolume, "rb")

	codeFile, err := os.Create(codeFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer codeFile.Close()

	_, err = codeFile.WriteString(sandbox.Code)
	if err != nil {
		log.Fatal(err)
	}

	// writing the submitted test cases into a file with name test_cases.txt
	testCasesFilePath := fmt.Sprintf("%s/test_cases.%s", requestVolume, "txt")

	testCasesFile, err := os.Create(testCasesFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer testCasesFile.Close()

	_, err = testCasesFile.WriteString(sandbox.TestCases)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	// to be implemented
}
