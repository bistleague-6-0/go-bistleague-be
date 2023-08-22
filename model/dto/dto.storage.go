package dto

type UploadStorageRequest struct {
	FileName   string `json:"file_name"`
	Base64Data string `json:"base64_data"`
}

type UploadStorageResponse struct {
	FileName string `json:"file_name"`
	Url      string `json: "url"`
}
