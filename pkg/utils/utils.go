package utils

import (
    "encoding/json"
    "io"
    "net/http"
    "runtime"
    "path/filepath"
)

func ParseBody(r *http.Request, x interface{}) error {
    defer r.Body.Close()

    body, err := io.ReadAll(r.Body)
    if err != nil {
        return err
    }

    if err := json.Unmarshal(body, x); err != nil {
        return err
    }

    return nil
}

// GetTemplateFilePath returns the absolute file path for the given template filename.
func GetTemplateFilePath(filename string) (string, error) {
    _, b, _, _ := runtime.Caller(0)
    basePath := filepath.Join(filepath.Dir(b), "../..")
    return filepath.Join(basePath, "templates", filename), nil
}

