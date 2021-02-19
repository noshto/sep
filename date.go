package sep

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

// Date represents a date without time component
type Date time.Time

// MarshalXMLAttr allows to generate "2006-01-02" format instead of RFC3339
func (d Date) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: time.Time(d).Format("2006-01-02"),
	}, nil
}

// UnmarshalXMLAttr unmarshals Date
func (d *Date) UnmarshalXMLAttr(attr xml.Attr) error {
	t, err := time.Parse("2006-01-02", attr.Value)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// UnmarshalJSON conformance
func (d *Date) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", tmp)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// MarshalJSON conformance
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

// DateTime represents a date with time component
type DateTime time.Time

// MarshalText renders conformance to Marshal protocol
func (d DateTime) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).Format(time.RFC3339)), nil
}

// UnmarshalText renders conformance to Unmarshal protocol
func (d *DateTime) UnmarshalText(text []byte) error {
	time, err := time.Parse(time.RFC3339, string(text))
	if err != nil {
		return err
	}
	*d = DateTime(time)
	return nil
}
