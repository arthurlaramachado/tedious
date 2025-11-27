package utils

import (
	"encoding/json"
	"os"

	mod "github.com/arthurlaramachado/tedious/internal/models"
)

const filepath = "./storage.json"

func SaveState(tasks []mod.Task) (string, bool) {
	jsonBlob, err := json.Marshal(tasks)

	if err != nil {
		return "Error while encoding", false
	}

	err_write := os.WriteFile(filepath, jsonBlob, 0644)

	if err_write != nil {
		return "Error while writing to file", false
	}

	return "Success", true
}

func ReadState() []mod.Task {
	var tasks []mod.Task

	jsonBlob, err := os.ReadFile(filepath)

	if err != nil {
		return nil
	}

	err = json.Unmarshal(jsonBlob, &tasks)

	if err != nil {
		return nil
	}

	return tasks
}
