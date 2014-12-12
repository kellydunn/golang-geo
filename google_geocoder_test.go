package geo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func GetMockResponse(s string) ([]byte, error) {
	dataPath := path.Join(s)
	_, readErr := os.Stat(dataPath)
	if readErr != nil && os.IsNotExist(readErr) {
		return nil, readErr
	}

	handler, handlerErr := os.Open(dataPath)
	if handlerErr != nil {
		return nil, handlerErr
	}

	data, readErr := ioutil.ReadAll(handler)
	if readErr != nil {
		return nil, readErr
	}

	return data, nil
}
