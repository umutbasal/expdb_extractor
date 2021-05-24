package extract

import (
	"encoding/xml"
)

// Element
type Cve struct {
	XMLName xml.Name `xml:"cve"`

	SchemaVersion string `xml:"schemaVersion,attr,omitempty"`

	Item []ItemType `xml:",any"`
}

// Element
type Comment struct {
	XMLName xml.Name `xml:"comment"`

	Voter string `xml:"voter,attr"`

	Text string `xml:",chardata"`
}

// Element
type Accept struct {
	XMLName xml.Name `xml:"accept"`

	Count string `xml:"count,attr"`

	Text string `xml:",chardata"`
}

// Element
type Modify struct {
	XMLName xml.Name `xml:"modify"`

	Count string `xml:"count,attr"`

	Text string `xml:",chardata"`
}

// Element
type Noop struct {
	XMLName xml.Name `xml:"noop"`

	Count string `xml:"count,attr"`

	Text string `xml:",chardata"`
}

// Element
type Recast struct {
	XMLName xml.Name `xml:"recast"`

	Count string `xml:"count,attr"`

	Text string `xml:",chardata"`
}

// Element
type Reject struct {
	XMLName xml.Name `xml:"reject"`

	Count string `xml:"count,attr"`

	Text string `xml:",chardata"`
}

// Element
type Reviewing struct {
	XMLName xml.Name `xml:"reviewing"`

	Count string `xml:"count,attr"`

	Text string `xml:",chardata"`
}

// Element
type Revote struct {
	XMLName xml.Name `xml:"revote"`

	Count string `xml:"count,attr"`

	Text string `xml:",chardata"`
}

// XSD ComplexType declarations

type ItemType struct {
	XMLName xml.Name

	Type TypeEnumType `xml:"type,attr"`

	Name string `xml:"name,attr"`

	Seq string `xml:"seq,attr"`

	Status string `xml:"status"`

	Phase *SpecificPhaseType `xml:"phase"`

	Desc string `xml:"desc"`

	Refs RefsType `xml:"refs"`

	Votes *VotesType `xml:"votes"`

	Comments *CommentsType `xml:"comments"`

	InnerXml string `xml:",innerxml"`
}

type CommentsType struct {
	XMLName xml.Name

	Comment []Comment `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type VotesType struct {
	XMLName xml.Name

	Accept *Accept `xml:"accept"`

	Modify *Modify `xml:"modify"`

	Noop *Noop `xml:"noop"`

	Recast *Recast `xml:"recast"`

	Reject *Reject `xml:"reject"`

	Reviewing *Reviewing `xml:"reviewing"`

	Revote *Revote `xml:"revote"`

	InnerXml string `xml:",innerxml"`
}

type SpecificPhaseType struct {
	XMLName xml.Name

	Date string `xml:"date,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type RefsType struct {
	XMLName xml.Name

	Ref []RefType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type RefType struct {
	XMLName xml.Name

	Source string `xml:"source,attr"`

	Url string `xml:"url,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type TypeEnumType string

const TypeEnumTypeCan TypeEnumType = "CAN"

const TypeEnumTypeCve TypeEnumType = "CVE"

type StatusEnumType string

const StatusEnumTypeEntry StatusEnumType = "Entry"

const StatusEnumTypeCandidate StatusEnumType = "Candidate"

type SimplePhaseEnumType string

const SimplePhaseEnumTypeProposed SimplePhaseEnumType = "Proposed"

const SimplePhaseEnumTypeInterim SimplePhaseEnumType = "Interim"

const SimplePhaseEnumTypeModified SimplePhaseEnumType = "Modified"

const SimplePhaseEnumTypeAssigned SimplePhaseEnumType = "Assigned"
