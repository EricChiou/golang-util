package readenv

import (
	"io/ioutil"
	"strings"
)

func LoadFile(filePath string) (map[string]string, error) {
	data := make(map[string]string)

	lines, err := fileHandler(filePath)
	if err != nil {
		return data, err
	}

	for _, line := range lines {
		val := strings.Split(line, "=")
		if len(val) == 2 {
			data[strings.TrimSpace(val[0])] = strings.TrimSpace(val[1])
		}
	}

	return data, nil
}

func fileHandler(filePath string) ([]string, error) {
	var lines []string
	binary, err := ioutil.ReadFile(filePath)
	if err != nil {
		return lines, err
	}

	str := string(binary)
	str = strings.Replace(str, "\r", "", -1)
	lines = strings.Split(str, "\n")
	return lines, nil
}
