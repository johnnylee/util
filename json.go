package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// JsonUnmarshal: attempt to unmarshal the json file into the given object.
func JsonUnmarshal(path string, v interface{}) error {
	var err error

	if path, err = ExpandPath(path); err != nil {
		return err
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, &v)
}

// JsonMarshal: attempt to write out the given object as json into the given
// file. This function won't overwrite an existing file.
func JsonMarshal(path string, v interface{}) error {
	var err error

	if path, err = ExpandPath(path); err != nil {
		return err
	}

	if FileExists(path) {
		return fmt.Errorf("Won't overwrite file: %v", path)
	}

	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err = json.Indent(&out, buf, "", "\t"); err != nil {
		return err
	}

	return ioutil.WriteFile(path, out.Bytes(), 0600)
}
