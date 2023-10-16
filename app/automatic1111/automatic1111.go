package automatic1111

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/linkisensei/civitdownloader/civit"
	"github.com/linkisensei/civitdownloader/util"
)

var instalationPath string = "F:\\stable-diffusion-webui"

type LoraConfigJson struct {
	SDVersion       string `json="sd version"`
	Description     string `json="description"`
	ActivationText  string `json="activation text"`
	PreferredWeight string `json="preferred weight"`
	Notes           string `json="notes"`
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
		// case "LOCON":
		// 	err = "models/LyCORIS"
	}

	if err != nil {
		return err
	}

	return nil
}

func CreateLoraConfigJson(basePath string, model *civit.CivitAIModel, version *civit.CivitAIModelVersion) error {
	var config LoraConfigJson

	// PAREI AQUI!

	return nil
}
