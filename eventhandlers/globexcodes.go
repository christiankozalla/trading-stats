package eventhandlers

import (
	_ "embed"
	"encoding/json"
)

//go:embed globexcodes.json
var jsonData []byte

var globexCodeToMultiplierMap = map[string]int{}

func init() {
	if err := json.Unmarshal(jsonData, &globexCodeToMultiplierMap); err != nil {
		panic("Error unmarshalling globex codes JSON: " + err.Error())
	}
}
