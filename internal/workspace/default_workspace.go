package workspace

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	WorkspaceGID  string `json:"workspace_gid"`
	WorkspaceName string `json:"workspace_name"`
}

func getConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user config dir: %w", err)
	}

	configPath := filepath.Join(configDir, "alfie")
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create config dir: %w", err)
	}

	return filepath.Join(configPath, "alfie_config.json"), nil
}

func SaveDefaultWorkspace(workspaceGID, workspaceName string) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	config := Config{
		WorkspaceGID:  workspaceGID,
		WorkspaceName: workspaceName,
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

func LoadDefaultWorkspace() (gid string, name string, loadError error) {
	path, err := getConfigFilePath()
	if err != nil {
		return "", "", err
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", "", errors.New("no default workspace set")
		}
		return "", "", fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return "", "", fmt.Errorf("failed to decode config file: %w", err)
	}

	return config.WorkspaceGID, config.WorkspaceName, nil
}
