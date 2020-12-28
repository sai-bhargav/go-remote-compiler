package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type statusResponse struct {
	Status string `json:"status"`
	Output string `json:"output"`
}

func ExecutionStatus(identifier string) statusResponse {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	res := statusResponse{}

	identifierPath := fmt.Sprintf("%s/payloads/%s", pwd, identifier)

	identifierExists, err := exists(identifierPath)
	if err != nil {
		log.Fatal(err)
	}

	if !identifierExists {
		res = statusResponse{"Code not submitted", ""}
		return res
	}

	res = statusResponse{"Running", ""}

	outputFilePath := fmt.Sprintf("%s/payloads/%s/output.txt", pwd, identifier)

	outputExists, err := exists(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if outputExists {
		output, err := ioutil.ReadFile(outputFilePath)
		if err != nil {
			log.Fatal(err)
		}

		res = statusResponse{"success", string(output)}
	}

	return res
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
