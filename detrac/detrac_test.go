package detrac_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"testing"

	"github.com/leeychee/annos/detrac"
)

var detracAnnoFilepath = "./MVI_20011.xml"

func TestUnmarshal(t *testing.T) {
	f, err := os.Open(detracAnnoFilepath)
	if err != nil {
		t.Errorf("Fail to open detrac annotation file: %s", detracAnnoFilepath)
	}
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, f)
	if err != nil {
		t.Errorf("Fail to read bytes from annotation file: %s", detracAnnoFilepath)
	}
	s := &detrac.Sequence{}
	err = detrac.Unmarshal(buf.Bytes(), s)
	if err != nil || len(s.Frames) != 664 {
		data, _ := detrac.Marshal(s)
		t.Logf("%s", data)
		t.Errorf("Fail to unmarshal sequence from file: %s", detracAnnoFilepath)
	} else {
		t.Logf("Anno %q has %d boxes, %d Frames.", s.Name, len(s.IgnoredRegion.Regions), len(s.Frames))
	}
}

func TestMarshal(t *testing.T) {
	s := detrac.Sequence{
		Name: "testname",
		Attribute: detrac.SeqAttribute{
			CameraState:  "unstable",
			SenceWeather: "sunny",
		},
		IgnoredRegion: detrac.IgnoredRegion{
			Regions: []detrac.Box{
				detrac.Box{
					Left:   0,
					Top:    1,
					Width:  2,
					Height: 1,
				},
				detrac.Box{
					Left:   0,
					Top:    1,
					Width:  2,
					Height: 1,
				},
			},
		},
		Frames: []detrac.Frame{
			detrac.Frame{
				Density: 1,
				Num:     1,
				Targets: detrac.Targets{
					Targets: []detrac.Target{
						detrac.Target{
							ID: 1,
							Box: detrac.Box{
								Left:   0,
								Top:    1,
								Width:  1,
								Height: 1,
							},
							Attribute: detrac.Attribute{
								Orientation:      0,
								Speed:            1,
								TrajectoryLength: 1,
								TruncationRatio:  1,
								VehicleType:      "car",
							},
							Occlusion: detrac.Occlusion{
								Occlusion: detrac.Overlap{
									ID:     1,
									Status: 1,
									Left:   1,
									Top:    1,
									Width:  1,
									Height: 1,
								},
							},
						},
					},
				},
			},
		},
	}
	b, err := detrac.Marshal(&s)
	if err != nil {
		t.Errorf("Fail to marshal %v", s)
	} else {
		t.Logf("%s%s", xml.Header, b)
	}
}
