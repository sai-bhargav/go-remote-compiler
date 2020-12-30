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
	Language      string `json:"lang"`
	Code          string `json:"code"`
	TestCases     string `json:"tc"`
	Identifier    string `json:"id"`
	Command       string
	FileExtension string
}

func Run(sandbox Sandbox) string {

	prepare(&sandbox)
	out := execute(sandbox)
	postProcessor(out, sandbox)

	return out
}

func prepare(sandbox *Sandbox) {

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

	sandbox.FileExtension = languages[sandbox.Language]["file_extension"]

	createVolume(sandbox)
}

func execute(sandbox Sandbox) string {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	sourceVolume := fmt.Sprintf("%s/payloads/%s", pwd, sandbox.Identifier)
	dockerCommand := fmt.Sprintf("docker run --rm -v %s:/app remote-code-compiler bash -c \"%s\"", sourceVolume, sandbox.Command)

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
	codeFilePath := fmt.Sprintf("%s/code.%s", requestVolume, sandbox.FileExtension)

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

func postProcessor(output string, sandbox Sandbox) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	requestVolume := fmt.Sprintf("%s/payloads/%s", pwd, sandbox.Identifier)

	codeFilePath := fmt.Sprintf("%s/output.%s", requestVolume, "txt")

	codeFile, err := os.Create(codeFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer codeFile.Close()

	_, err = codeFile.WriteString(output)
	if err != nil {
		log.Fatal(err)
	}

}
