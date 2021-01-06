package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func JsonDecode(readCloser io.ReadCloser, v interface{}) error {
	reader, err := ioutil.ReadAll(readCloser); if err != nil {
		return err
	}

	err = json.Unmarshal(reader, v); if err != nil {
		return err
	}

	return nil
}
