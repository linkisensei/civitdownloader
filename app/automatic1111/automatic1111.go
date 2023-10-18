package automatic1111

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/linkisensei/civitdownloader/app/config"
	"github.com/linkisensei/civitdownloader/civit"
	"github.com/linkisensei/civitdownloader/util"
)

type LoraConfigJson struct {
	SDVersion       string  `json="sd version"`
	Description     string  `json="description"`
	ActivationText  string  `json="activation text"`
	PreferredWeight float64 `json="preferred weight"`
	Notes           string  `json="notes"`
}

type LycorisConfigJson struct {
	Description string `json="description"`
	Notes       string `json="notes"`
}

func GetModelPathFromCivitAiModel(model *civit.CivitAIModel) (string, error) {
	var relativePath string

	switch strings.ToUpper(model.Type) {
	case "LORA":
		relativePath = "models/Lora"
	case "LOCON":
		relativePath = "models/LyCORIS"
	case "TEXTUALINVERSION":
		relativePath = "embeddings"
	case "HYPERNETWORK":
		relativePath = "models/hypernetworks"
	case "CHECKPOINT":
		relativePath = "models/Stable-diffusion"
	}

	if relativePath == "" {
		return relativePath, errors.New(fmt.Sprintf("unknow model type %s", model.Type))
	}

	instalationPath := config.Config.GetString(config.INSTALLATION_PATH)
	filePath := filepath.Join(instalationPath, relativePath)

	// Making sure that
	err := util.CreateDirectoryIfDoesntExists(filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func CreateConfigJson(basePath string, model *civit.CivitAIModel, version *civit.CivitAIModelVersion) error {
	var err error

	switch strings.ToUpper(model.Type) {
	case "LORA":
		err = CreateLoraConfigJson(basePath, model, version)
	case "LOCON":
		err = CreateLycorisConfigJson(basePath, model, version)
	}

	if err != nil {
		return err
	}

	return nil
}

func CreateLoraConfigJson(basePath string, model *civit.CivitAIModel, version *civit.CivitAIModelVersion) error {
	config := LoraConfigJson{
		Description:     fmt.Sprintf("https://civitai.com/models/%d", model.Id),
		PreferredWeight: 0.8,
	}

	if len(version.TrainedWords) > 0 {
		config.ActivationText = strings.Join(version.TrainedWords, ", ")
	}

	baseModelVersion, _ := strconv.ParseFloat(version.BaseModel, 32)
	if baseModelVersion < 2 {
		config.SDVersion = "SD1"
	}

	// fmt.Printf("%+v", config)

	encodedConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	modelFilename, _ := version.GetFilename()
	jsonFilePath := strings.TrimSuffix(modelFilename, filepath.Ext(modelFilename)) + ".json"
	jsonFilePath = filepath.Join(basePath, jsonFilePath)

	// Open the file for writing
	err = ioutil.WriteFile(jsonFilePath, encodedConfig, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func CreateLycorisConfigJson(basePath string, model *civit.CivitAIModel, version *civit.CivitAIModelVersion) error {
	config := LycorisConfigJson{
		Description: fmt.Sprintf("https://civitai.com/models/%d", model.Id),
	}

	encodedConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	modelFilename, _ := version.GetFilename()
	jsonFilePath := strings.TrimSuffix(modelFilename, filepath.Ext(modelFilename)) + ".json"
	jsonFilePath = filepath.Join(basePath, jsonFilePath)

	// Open the file for writing
	err = ioutil.WriteFile(jsonFilePath, encodedConfig, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
