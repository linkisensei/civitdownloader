package app

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/linkisensei/civitdownloader/app/automatic1111"
	"github.com/linkisensei/civitdownloader/civit"
	"github.com/linkisensei/civitdownloader/util"
)

func DownloadModel(modelUrl string) error {
	requestInfo, err := civit.CreateRequestInfoFromUrl(modelUrl)
	if err != nil {
		return err
	}

	// Retrieving Model
	civitModel, err := civit.GetModel(requestInfo.ModelId)
	if err != nil {
		return err
	}

	greenColor := color.New(color.FgGreen).Add(color.Underline)
	greenColor.Printf("\n\nModel: %s", civitModel.Name+"\n")

	// Extracting Model Version and Model File
	modelVersion := civitModel.GetVersion(requestInfo.VersionId)
	modelFile, err := modelVersion.GetModelFile()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	greenColor.Printf("Version: %s", modelVersion.Name+"\n\n")

	downloadUrl := modelFile.DownloadUrl
	filename := modelFile.Name

	if downloadUrl == "" || filename == "" {
		return errors.New("invalid model download url or filename")
	}

	basePath, err := automatic1111.GetModelPathFromCivitAiModel(&civitModel)
	if err != nil {
		return err
	}

	filePath := filepath.Join(basePath, filename)

	// Downloading the Model Version File
	greenColor.Printf("1) Model File:\n")
	util.DownloadFile(downloadUrl, filePath)

	// Downloading a Model Version Image
	modelImage := modelVersion.GetRandomModelImage()
	if modelImage != nil {
		imageFilename := modelImage.GetFilenameForModelVersion(modelVersion)
		imagePath := filepath.Join(basePath, imageFilename)
		greenColor.Printf("2) Image File:\n")
		util.DownloadFile(modelImage.Url, imagePath)
	}

	// Generating Model Version Config JSON
	greenColor.Printf("3) Config JSON:\n")
	err = automatic1111.CreateConfigJson(basePath, &civitModel, modelVersion)
	if err != nil {
		fmt.Printf("Skipped")
	} else {
		fmt.Printf("Generated")
	}

	return nil
}
