package ctypes

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// Custom Tab Stop
type Tab struct {
	// Tab Stop Type
	Val *stypes.CustTabStop `xml:"val,attr,omitempty"`

	//Tab Stop Position
	Position *int `xml:"pos,attr,omitempty"`

	//Custom Tab Stop Leader Character
	LeaderChar *stypes.CustLeadChar `xml:"leader,attr,omitempty"`
}

func (t *Tab) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tab"
	start.Attr = []xml.Attr{}

	if t.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"}, Value: string(*t.Val)})
	}

	if t.Position != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pos"}, Value: strconv.Itoa(*t.Position)})
	}

	if t.LeaderChar != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "leader"}, Value: string(*t.LeaderChar)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type Tabs struct {
	Tab []Tab `xml:"tab,omitempty"`
}

func (t Tabs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if len(t.Tab) == 0 {
		return nil
	}

	// Create the enclosing XML element
	start.Name = xml.Name{Local: "w:tabs"}

	err := e.EncodeToken(start)
	if err != nil {
		return fmt.Errorf("error encoding start element: %v", err)
	}

	for _, tab := range t.Tab {
		if err := tab.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("error encoding tab: %v", err)
		}
	}

	err = e.EncodeToken(start.End())
	if err != nil {
		return fmt.Errorf("error encoding end element: %v", err)
	}

	return nil
}