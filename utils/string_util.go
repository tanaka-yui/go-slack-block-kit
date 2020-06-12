package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJson(obj interface{}) string {
	var buf bytes.Buffer
	b, _ := json.Marshal(obj)
	buf.Write(b)
	return buf.String()
}

func ToString(io io.Reader) string {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(io); err != nil {
		return ""
	}
	return buf.String()
}
