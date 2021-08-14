package tools

import (
	"encoding/csv"
	"os"

	"github.com/attreios/holmes"
	"gopkg.in/yaml.v2"
)

func ReadYamlFile(url string, result interface{}) {
	h := holmes.New(os.Getenv("LOG_MODE"), "utils.ReadYaml")
	f, err := os.Open(url)
	if err != nil {
		h.FatalError("Could not open file: %v", err)
	}
	defer f.Close()
	d := yaml.NewDecoder(f)
	if err := d.Decode(result); err != nil {
		h.FatalError("Could not decode YAML: %v", err)
	}
}

func ReadCSVFile(url string) [][]string {
	h := holmes.New(os.Getenv("LOG_MODE"), "utils.ReadCSV")

	f, err := os.Open(url)
	if err != nil {
		h.FatalError("Could not open file: %v", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	result, err := r.ReadAll()
	if err != nil {
		h.FatalError("Could not read CSV: %v", err)
	}
	return result
}
