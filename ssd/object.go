package ssd

import (
	"encoding/xml"
)

type Annotation struct {
	XMLName  xml.Name `xml:"annotation"`
	Folder   string   `xml:"folder"`
	Filename string   `xml:"filename"`
	Source   Source   `xml:"source"`
	Owner    Owner    `xml:"owner"`
	Size     Size     `xml:"size"`
	Objects  []Object `xml:"object"`
}

type Source struct {
	Database   string `xml:"database"`
	Annotation string `xml:"annotation"`
	Image      string `xml:"image"`
	Flickrid   string `xml:"flickrid"`
}

type Owner struct {
	Flickrid string `xml:"flickrid"`
	Name     string `xml:"name"`
}

type Size struct {
	Width  int `xml:"width"`
	Height int `xml:"height"`
	Depth  int `xml:"depth"`
}

type Object struct {
	Name      string `xml:"name"`
	Pose      string `xml:"pose"`
	Truncated int    `xml:"truncated"`
	Difficult int    `xml:"difficult"`
	Bndbox    Bndbox `xml:"bndbox"`
}

type Bndbox struct {
	Xmin int `xml:"xmin"`
	Ymin int `xml:"ymin"`
	Xmax int `xml:"xmax"`
	Ymax int `xml:"ymax"`
}
