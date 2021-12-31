package boards

import (
	"encoding/json"
)

func (b *Board) ArrayToJson(arr [][]int16) (string, error) {

	terrainString, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(terrainString), nil
}

func (b *Board) JsonToArray(arr string) ([][]int16, error) {

	var array [][]int16
	err := json.Unmarshal([]byte(arr), &array)
	if err != nil {
		return array, err
	}
	return array, nil
}
