package main

import (
	"encoding/xml"
	"os"
)

type TargetDef struct {
	Name           string      `xml:"name,attr"`
	SequenceNumber int         `xml:"sequenceNumber,attr"`
	Locations      []TargetLoc `xml:"locations>location"`
}

type TargetLoc struct {
	Type  string       `xml:"type,attr"`
	Repo  TargetRepo   `xml:"repository"`
	Units []TargetUnit `xml:"unit"`
}

type TargetRepo struct {
	URL string `xml:"location,attr"`
}

type TargetUnit struct {
	ID      string `xml:"id,attr"`
	Version string `xml:"version,attr"`
}

func parseTargetDefinition(file string) (*TargetDef, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	def := &TargetDef{}
	err = xml.Unmarshal(data, def)
	if err != nil {
		return nil, err
	}
	return def, nil
}
