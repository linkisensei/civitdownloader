package app

import (
	"errors"
	"fmt"
	"path/filepath"

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

	// Extracting Model Version and Model File
	modelVersion := civitModel.GetVersion(requestInfo.VersionId)
	modelFile, err := modelVersion.GetModelFile()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

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

	fmt.Println(downloadUrl)
	fmt.Println(filePath)

	// Downloading the Model Version File
	util.DownloadFile(downloadUrl, filePath)

	// Downloading a Model Version Image
	modelImage := modelVersion.GetRandomModelImage()
	if modelImage != nil {
		imageFilename := modelImage.GetFilenameForModelVersion(modelVersion)
		imagePath := filepath.Join(basePath, imageFilename)
		util.DownloadFile(modelImage.Url, imagePath)
	}

	// Generating Model Version Config JSON
	_ = automatic1111.CreateConfigJson(basePath, &civitModel, modelVersion)

	return nil
}
