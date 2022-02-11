package ser

import (
	"bytes"
	"encoding/gob"

	"github.com/pkg/errors"
)

func Serialize(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	err := gob.NewEncoder(&buffer).Encode(v)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"gob encode value %v",
			v,
		)
	}

	return buffer.Bytes(), nil
}

func Deserialize(b []byte) (interface{}, error) {
	decoder := gob.NewDecoder(bytes.NewBuffer(b))

	var buffer interface{}
	err := decoder.Decode(&buffer)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"gob decode",
		)
	}

	return buffer, nil
}
