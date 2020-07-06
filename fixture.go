package fixture

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Save a struct to a file with a path and a file name
// path can be an empty string to save to current directory.
// the json is saved in pretty print
func Save(inter interface{}, path string, fileName string) error {
	js, err := json.MarshalIndent(inter, "", "	")
	if err != nil {
		return err
	}

	filePath := filepath.Join(path, fileName)

	err = ioutil.WriteFile(filePath, js, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Read a struct from a file with a path and a file name
// path can be an empty string to save to current directory.
// the json file can be in any format pretty or no spaces.
func Read(inter interface{}, path string, fileName string) error {
	filePath := filepath.Join(path, fileName)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(byteValue), &inter)
	if err != nil {
		return err
	}

	return nil
}
