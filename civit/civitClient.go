package civit

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var httpClient *http.Client

const baseEndpoint = "https://civitai.com"

type CivitAIModelFile struct {
	Id          int64   `json:"id"`
	Size        float64 `json:"sizeKB"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	DownloadUrl string  `json:"downloadUrl"`
}

type CivitAIImage struct {
	Url string `json:"url"`
}

func (img *CivitAIImage) GetFilenameForModelVersion(version *CivitAIModelVersion) string {
	modelFilename, err := version.GetFilename()
	if err != nil {
		return ""
	}

	fileExtension := filepath.Ext(img.Url)
	modelName := strings.TrimSuffix(modelFilename, filepath.Ext(modelFilename))

	newFilename := modelName + fileExtension
	return newFilename
}

type CivitAIModelVersion struct {
	Id           int64              `json:"id"`
	ModelId      int64              `json:"modelId"`
	Name         string             `json:"name"`
	TrainedWords []string           `json:"trainedWords"`
	BaseModel    string             `json:"baseModel"`
	Files        []CivitAIModelFile `json:"files"`
	Images       []CivitAIImage     `json:"images"`
}

type CivitAIModel struct {
	Id       int64                 `json:"id"`
	Name     string                `json:"name"`
	Type     string                `json:"type"`
	Versions []CivitAIModelVersion `json:"modelVersions"`
}

type modelRequestInfo struct {
	ModelId   int64
	VersionId int64
}

func (v *CivitAIModelVersion) GetModelFile() (*CivitAIModelFile, error) {
	var modelFile *CivitAIModelFile

	for _, file := range v.Files {
		if file.Type == "Model" {
			modelFile = &file
			break
		}
	}

	if modelFile == nil {
		return modelFile, errors.New("missing model filename")
	}

	return modelFile, nil
}

func (v *CivitAIModelVersion) GetFilename() (string, error) {
	modelFile, err := v.GetModelFile()
	if err != nil {
		return "", err
	}

	if modelFile.Name == "" {
		return "", errors.New("missing model filename")
	}
	return modelFile.Name, nil
}

func (v *CivitAIModelVersion) GetDownloadUrl() (string, error) {
	modelFile, err := v.GetModelFile()
	if err != nil {
		return "", err
	}

	if modelFile.DownloadUrl == "" {
		return "", errors.New("missing model download url")
	}
	return modelFile.DownloadUrl, nil
}

func (v *CivitAIModelVersion) GetRandomModelImage() *CivitAIImage {
	imageCount := len(v.Images)
	if imageCount == 0 {
		return &CivitAIImage{}
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(imageCount)
	randomItem := v.Images[randomIndex]
	return &randomItem
}

/**
 * If versionId is null, gets the latest version (index 0)
 */
func (m *CivitAIModel) GetVersion(versionId int64) *CivitAIModelVersion {
	var desiredVersion *CivitAIModelVersion

	if versionId == 0 {
		return &m.Versions[0]
	}

	for index, version := range m.Versions {
		if version.Id == versionId {
			desiredVersion = &m.Versions[index]
			break
		}
	}

	return desiredVersion
}

func getJson(url string, target interface{}) error {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 20 * time.Second}
	}
	response, err := httpClient.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

func GetModel(modelId int64) (CivitAIModel, error) {
	url := baseEndpoint + "/api/v1/models/" + strconv.FormatInt(modelId, 10)

	var model CivitAIModel

	err := getJson(url, &model)
	if err != nil {
		fmt.Printf("Error getting model: %s\n", err.Error())
		return CivitAIModel{}, err
	}

	return model, nil
}

func CreateRequestInfoFromUrl(modelUrl string) (modelRequestInfo, error) {

	var request modelRequestInfo

	parsedURL, err := url.Parse(modelUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return request, err
	}

	// Getting Model ID
	pattern := `/models/(\d+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(parsedURL.Path)

	if len(matches) == 0 {
		return request, errors.New("invalid model url (missing model id)")

	} else {
		idStr := matches[1]
		modelId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return request, errors.New("invalid model id")
		}
		request.ModelId = modelId
	}

	// Getting Model Version, if specified
	modelVersionId, err := strconv.ParseInt(parsedURL.Query().Get("modelVersionId"), 10, 64)
	if err == nil {
		request.VersionId = modelVersionId
	}

	return request, nil
}
