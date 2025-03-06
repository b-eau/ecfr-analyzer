package write_json

import (
	"encoding/json"
	"os"
)

func WriteJSONToFile(fileName string, content any) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(content)
}
