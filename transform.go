package annos

import (
	"errors"
	"fmt"

	"github.com/leeychee/annos/detrac"
	"github.com/leeychee/annos/ssd"
)

// Detrac2Ssd transfer detrac to ssd
func Detrac2Ssd(d *detrac.Sequence) ([]*ssd.Annotation, error) {
	if d == nil {
		return nil, errors.New("invalid parameter, d should not nil")
	}
	ss := make([]*ssd.Annotation, len(d.Frames))
	for i, f := range d.Frames {
		objs := make([]ssd.Object, len(f.Targets.Targets))
		for i, t := range f.Targets.Targets {
			obj := ssd.Object{
				Name:      t.Attribute.VehicleType,
				Pose:      "Unspecified",
				Truncated: 1,
				Difficult: 0,
				Bndbox:    detracBox2SsdBndbox(t.Box),
			}
			objs[i] = obj
		}
		s := &ssd.Annotation{
			Folder:   d.Name,
			Filename: fmt.Sprintf("%s_img%05d.jpg", d.Name, f.Num),
			Size: ssd.Size{
				Width:  detrac.ImageWidth,
				Height: detrac.ImageHeight,
				Depth:  detrac.ImageDepth,
			},
			Source: ssd.Source{
				Database:   "DETRAC",
				Annotation: "DETRAC",
				Image:      "detrac",
				Flickrid:   "detrac",
			},
			Owner: ssd.Owner{
				Flickrid: "detrac",
				Name:     "detrac",
			},
			Objects: objs,
		}
		ss[i] = s
	}
	return ss, nil
}

func detracBox2SsdBndbox(d detrac.Box) ssd.Bndbox {
	return ssd.Bndbox{
		Xmin: int(d.Left + 0.5),
		Ymin: int(d.Top + 0.5),
		Xmax: int(d.Left + d.Width + 0.5),
		Ymax: int(d.Top + d.Height + 0.5),
	}
}
