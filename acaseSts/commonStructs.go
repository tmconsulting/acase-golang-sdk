package acaseSts

import (
	"encoding/xml"
)

type LanguageTypeEnum string
const (
	Ru LanguageTypeEnum = "ru"
	En LanguageTypeEnum = "en"
)

type AdmUnitActionNameEnum string
const (
	List AdmUnitActionNameEnum = "LIST"
)

type ErrorType struct {
	XMLName		xml.Name		`xml:"Error"`
	Type		ErrorTypeType	`xml:"Type"`
	Pointer		string			`xml:"Pointer"`
	Code		string			`xml:"Code,attr"`
	Description	string			`xml:"Description,attr"`
}

type ErrorTypeType struct {
	XMLName		xml.Name	`xml:"Type"`
	Code		string		`xml:"Code,attr"`
	Name		string		`xml:"Name,attr"`
}

type CountryType struct {
	XMLName		xml.Name		`xml:"Country"`
	Code		string			`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Position	PositionType	`xml:"Position"`
}

type TypeOfPlaceType struct {
	XMLName	xml.Name	`xml:"TypeOfPlace"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}

type PositionType struct {
	XMLName		xml.Name	`xml:"Position"`
	Latitude	string		`xml:"Latitude,attr"`
	Longitude	string		`xml:"Longitude,attr"`
}
