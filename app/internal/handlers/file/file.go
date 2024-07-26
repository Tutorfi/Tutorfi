package filehandler

import (
	"app/internal/storage"

	"github.com/labstack/echo/v4"
)

type FileHandler struct {
	store storage.Storage
    fileServerHost string
}

func New(store storage.Storage, fileServerHost string) *FileHandler {
    return &FileHandler{
        store: store,
        fileServerHost: fileServerHost,
    }
}

// filename and path in string

func (handle *FileHandler) GetFile(c echo.Context) error {

    return nil;
}
