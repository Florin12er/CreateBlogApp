package utils

import (
    "encoding/json"
    "io"
    "net/http"
    "runtime"
    "path/filepath"
)

func ParseBody(r *http.Request, x interface{}) {
    if body , err := io.ReadAll(r.Body); err == nil {
        if err := json.Unmarshal([]byte(body), x); err != nil {
            return
        }
    }
}
func GetTemplatePath(filename string) (string, error) {
    _, b, _, _ := runtime.Caller(0)
    basePath := filepath.Join(filepath.Dir(b), "../..")
    return filepath.Join(basePath, "templates", filename), nil
}

