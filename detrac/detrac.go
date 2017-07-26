package detrac

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

const (
	// ImageWidth image width in DETRAC dataset
	ImageWidth = 960
	// ImageHeight image height in DETRAC dataset
	ImageHeight = 540
	// ImageDepth image depth in DETRAC dataset
	ImageDepth = 3
)

// ReadFromFile read sequence from file
func ReadFromFile(filename string) (*Sequence, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("fail to read sequence from %s, %s", filename, err)
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, f)
	if err != nil {
		return nil, fmt.Errorf("fail to read sequence from %s, %s", filename, err)
	}

	s := &Sequence{}
	err = Unmarshal(buf.Bytes(), s)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal sequence, %s", err)
	}
	return s, nil
}

// Marshal marshal sequence by the default indent
func Marshal(v *Sequence) ([]byte, error) {
	return MarshalIndent(v, "", "  ")
}

// MarshalIndent marshal sequence by the given indent
func MarshalIndent(v *Sequence, prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(v, prefix, indent)
}

// Unmarshal unmarshal data to v
func Unmarshal(data []byte, v *Sequence) error {
	return xml.Unmarshal(data, v)
}
