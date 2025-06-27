package utils

import (
	"encoding/json"
	"os"

	"jerobas.com/yepee/types"
)

func ReadJSON(path string) (types.RoutesType, error) {
	var localRoutes types.RoutesType
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &localRoutes)
	if err != nil {
		return nil, err
	}

	return localRoutes, nil
}

func UpdateJSON(path string, newJson types.RoutesType) error {
	marshaledJSON, err := json.Marshal(newJson)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, marshaledJSON, 0064)
	if err != nil {
		return err
	}

	return nil
}
