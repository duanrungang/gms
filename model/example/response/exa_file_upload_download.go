package response

import (
	"gms/model/example"
)

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
