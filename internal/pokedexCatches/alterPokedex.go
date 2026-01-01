package pokedexCatches

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"github.com/Bention99/pokedexcli/internal/api"
)

func SaveCaughtJSON(path string, caught map[string]api.Pokemon) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(caught, "", "  ")
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func LoadCaughtJSON(path string) (map[string]api.Pokemon, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return make(map[string]api.Pokemon), nil
		}
		return nil, err
	}

	var caught map[string]api.Pokemon
	if err := json.Unmarshal(data, &caught); err != nil {
		return nil, err
	}
	if caught == nil {
		caught = make(map[string]api.Pokemon)
	}
	return caught, nil
}

func DeleteCaughtFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return nil
}
