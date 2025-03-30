package gopickle

import (
	"encoding/gob"
	"os"
)

// Dump serializes the given data and writes it to a file.
func Dump(filename string, data any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return file.Close()
}

// Load deserializes data from a file and stores it in the provided variable.
func Load(filename string, data any) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		return err
	}
	return file.Close()
}
