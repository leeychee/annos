package detrac

import "encoding/xml"

// Sequence DETRAC datasets annotation struct.
type Sequence struct {
	XMLName       xml.Name      `xml:"sequence"`
	Name          string        `xml:"name,attr"`
	Attribute     SeqAttribute  `xml:"sequence_attribute"`
	IgnoredRegion IgnoredRegion `xml:"ignored_region"`
	Frames        []Frame       `xml:"frame"`
}

type SeqAttribute struct {
	CameraState  string `xml:"camera_state,attr"`
	SenceWeather string `xml:"sence_weather,attr"`
}

type IgnoredRegion struct {
	Regions []Box `xml:"box"`
}

type Box struct {
	Left   float32 `xml:"left,attr"`
	Top    float32 `xml:"top,attr"`
	Width  float32 `xml:"width,attr"`
	Height float32 `xml:"height,attr"`
}

type Frame struct {
	Density int     `xml:"density,attr"`
	Num     int     `xml:"num,attr"`
	Targets Targets `xml:"target_list"`
}

type Targets struct {
	Targets []Target `xml:"target"`
}

type Target struct {
	ID        int       `xml:"id,attr"`
	Box       Box       `xml:"box"`
	Attribute Attribute `xml:"attribute"`
	Occlusion Occlusion `xml:"occlusion,omitempty"`
}

type Attribute struct {
	Orientation      float32 `xml:"orientation,attr"`
	Speed            float32 `xml:"speed,attr"`
	TrajectoryLength int     `xml:"trajectory_length,attr"`
	TruncationRatio  int     `xml:"truncation_ration,attr"`
	VehicleType      string  `xml:"vehicle_type,attr"`
}

type Occlusion struct {
	Occlusion Overlap `xml:"region_overlap"`
}

type Overlap struct {
	ID     int     `xml:"occlusion_id,attr"`
	Status int     `xml:"occlusion_status,attr"`
	Left   float32 `xml:"left,attr"`
	Top    float32 `xml:"top,attr"`
	Width  float32 `xml:"width,attr"`
	Height float32 `xml:"height,attr"`
}
