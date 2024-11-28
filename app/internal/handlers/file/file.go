package filehandler

import (
	"app/internal/storage"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FileHandler struct {
	store          storage.Storage
	fileServerHost string
}

func New(store storage.Storage, fileServerHost string) *FileHandler {
	return &FileHandler{
		store:          store,
		fileServerHost: fileServerHost,
	}
}

// filename and path in string

// Make sure to only respond with the file data
func (handle *FileHandler) GetFile(c echo.Context) error {
	res, err := http.Get(handle.fileServerHost + "/test/path/data.txt")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "File server is down")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "File server is down")
	}
	fmt.Println(string(body))
	return nil
}

func (handle *FileHandler) SaveFile(c echo.Context) error {
	data := []byte("some data")
	responseBody := bytes.NewBuffer(data)
	res, err := http.NewRequest(http.MethodGet, handle.fileServerHost+"/test/path/data.txt", responseBody)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "File server is down")
	}
    defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "File server is down")
	}
	fmt.Println(string(body))
	return nil
}
