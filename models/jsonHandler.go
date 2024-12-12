package models

import (
	"encoding/json"
	"io"
	"os"
)

type Link struct {
	Link     string `json:"link"`
	Image    string `json:"image"`
	ImageAlt string `json:"image_alt"`
}

type AboutMe struct {
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Job   string `json:"job"`
	About string `json:"about"`
	Links []Link `json:"links"`
}

type Education struct {
	Icon   string `json:"icon"`
	School string `json:"school"`
	Degree string `json:"degree"`
}

type Job struct {
	Logo             string   `json:"logo"`
	Name             string   `json:"name"`
	Position         string   `json:"position"`
	StartDate        string   `json:"startDate"`
	EndDate          string   `json:"endDate"`
	Responsibilities []string `json:"responsibilities"`
}

type Project struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	RepoLink string `json:"repoLink"`
	Target   string `json:"target"`
}

type MainContent struct {
	Education []Education `json:"education"`
	Jobs      []Job       `json:"jobs"`
	Projects  []Project   `json:"projects"`
}

func LoadJSON[T any](filePath string) (T, error) {
	var result T

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(byteValue, &result)
	return result, err
}
