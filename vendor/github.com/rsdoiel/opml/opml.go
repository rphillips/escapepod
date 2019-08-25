//
// Package opml provides basic utility functions for working with OPML files.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2019, R. S. Doiel
// All rights not granted herein are expressly reserved by R. S. Doiel.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package opml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const (
	// Version of the ompl package, useful for display in cli tools
	Version = `v0.0.6`

	// The license for the ompl package, useful for display in cli tools
	LicenseText = `
%s %s

Copyright (c) 2019, R. S. Doiel
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of opml nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`
)

type CustomAttrs []xml.Attr

func (cattr CustomAttrs) MarshalJSON() ([]byte, error) {
	m := map[string]string{}
	for _, attr := range cattr {
		k := attr.Name.Local
		v := attr.Value
		if k != "" {
			m[k] = v
		}
	}

	return json.Marshal(m)
}

// OPML is the root structure for holding an OPML document
type OPML struct {
	XMLName   xml.Name    `xml:"opml" json:"-"`
	Version   string      `xml:"version,attr" json:"version"`
	Head      *Head       `xml:"head" json:"head"`
	Body      *Body       `xml:"body" json:"body"`
	OtherAttr CustomAttrs `xml:",any,attr" json:"other_attrs,omitempty"`
}

// Head holds the metadata for an OPML document
type Head struct {
	XMLName         xml.Name    `json:"-"`
	Title           string      `xml:"title,omitempty" json:"title,omitempty"`
	Created         string      `xml:"dateCreated,omitempty" json:"dateCreated,omitempty"`   // RFC 882 date and time
	Modified        string      `xml:"dateModified,omitempty" json:"dataModified,omitempty"` // RFC 882 date and time
	OwnerName       string      `xml:"ownerName,omitempty" json:"ownerName,omitempty"`
	OwnerEmail      string      `xml:"ownerEmail,omitempty" json:"ownerEmail,omitempty"`
	OwnerID         string      `xml:"OwnerId,omitempty" json:"OwnerId,omitempty"`               // url
	Docs            string      `xml:"docs,omitempty" json:"docs,omitempty"`                     // url
	ExpansionState  string      `xml:"expansionState,omitempty" json:"expansionState,omitempty"` // array of numbers
	VertScrollState int         `xml:"vertScrollState,omitempty" json:"vertScrollState,omitempty"`
	WindowTop       int         `xml:"windowTop,omitempty" json:"windowTop,omitempty"`
	WindowLeft      int         `xml:"windowLeft,omitempty" json:"windowLeft,omitempty"`
	WindowBottom    int         `xml:"windowBottom,omitempty" json:"windowBottom,omitempty"`
	WindowRight     int         `xml:"windowRight,omitempty" json:"windowRight,omitempty"`
	OtherAttr       CustomAttrs `xml:",any,attr" json:"other_attrs,omitempty"`
}

// Body holds the outline for an OPML document
type Body struct {
	XMLName   xml.Name    `json:"-"`
	Outline   OutlineList `xml:"outline" json:"outline"`
	OtherAttr CustomAttrs `xml:",any,attr" json:"other_attrs,omitempty"`
}

// Outline is the primary element of an OPML document, may hold sub-Outlines
type Outline struct {
	XMLName      xml.Name    `json:"-"`
	Text         string      `xml:"text,attr" json:"text"`
	Type         string      `xml:"type,attr,omitempty" json:"type,omitempty"`
	Title        string      `xml:"title,attr,omitempty" json:"title,omitempty"`
	IsComment    bool        `xml:"isComment,attr,omitempty" json:"isComment,omitempty"`
	IsBreakpoint bool        `xml:"isBreakpoint,attr,omitempty" json:"isBreakpoint,omitempty"`
	Created      string      `xml:"created,attr,omitempty" json:"created,omitempty"` // RFC 882 date and time
	Category     string      `xml:"category,attr,omitempty" json:"category,omitempty"`
	XMLURL       string      `xml:"xmlUrl,attr,omitempty" json:"xmlUrl,omitempty"`   // url
	HTMLURL      string      `xml:"htmlUrl,attr,omitempty" json:"htmlUrl,omitempty"` // url
	Language     string      `xml:"langauge,attr,omitempty" json:"language,omitempty"`
	Description  string      `xml:"description,attr,omitempty" json:"description,omitempty"`
	Version      string      `xml:"version,attr,omitempty" json:"version,omitempty"`
	URL          string      `xml:"url,attr,omitempty" json:"url,omitempty"` // url
	Outline      OutlineList `xml:"outline,omitempty" json:"outline,omitempty"`
	OtherAttr    CustomAttrs `xml:",any,attr" json:"other_attrs,omitempty"`
}

type OutlineList []Outline
type ByText []Outline
type ByTextCaseInsensitive []Outline
type ByType []Outline
type ByTitle []Outline
type ByTitleCaseInsensitive []Outline

// New creates an empty OPML structure
func New() *OPML {
	o := new(OPML)
	// OPML spec support
	o.Version = `2.0`

	o.Head = new(Head)
	o.Body = new(Body)
	return o
}

func (h *Head) String() string {
	s, _ := xml.Marshal(h)
	return string(s)
}

func (b *Body) String() string {
	s, _ := xml.Marshal(b)
	return string(s)
}

func (ol *Outline) String() string {
	s, _ := xml.Marshal(ol)
	return string(s)
}

// HasChildren return true if the outline element has a populated child outline
func (ol *Outline) HasChildren() bool {
	if len(ol.Outline) > 0 {
		return true
	}
	return false
}

// Append adds one or more outline element to the end of an outline list
func (ol OutlineList) Append(elems ...*Outline) error {
	for _, elem := range elems {
		i := len(ol)
		ol = append(ol, *elem)
		if len(ol) != (i+1) || ol[i].Text != elem.Text {
			return fmt.Errorf("failed to append element")
		}
	}
	return nil
}

// Adds one or more children to outline element
func (ol *Outline) AppendChild(elems ...*Outline) error {
	for _, elem := range elems {
		err := ol.Outline.Append(elem)
		if err != nil {
			return fmt.Errorf("failed to add child, %s", err)
		}
	}
	return nil
}

// Append one or more Body.Outline lists to the current OPML structure
func (o *OPML) Append(outlines ...*OPML) error {
	i := len(o.Body.Outline)
	for _, next := range outlines {
		for _, elem := range next.Body.Outline {
			o.Body.Outline = append(o.Body.Outline, elem)
			i += 1
		}
	}
	if len(o.Body.Outline) != i {
		return fmt.Errorf("Failed to add all outline elements, exlected %d, have %d", i, len(o.Body.Outline))
	}
	return nil
}

func (o *OPML) String() string {
	if o.Body != nil {
		if o.Body.Outline == nil {
			o.Body.Outline = make(OutlineList, 1)
			o.Body.Outline.Append(&Outline{
				Text: "",
			})
		} else if len(o.Body.Outline) == 0 {
			o.Body.Outline.Append(&Outline{
				Text: "",
			})
		}
	}
	s, _ := xml.Marshal(o)
	return string(s)
}

// Len for ByText sort of Outline
func (a ByText) Len() int {
	return len(a)
}

// Swap for ByText sort of Outline
func (a ByText) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less for ByText sort of Outline
func (a ByText) Less(i, j int) bool {
	return strings.Compare(a[i].Text, a[j].Text) == -1
}

// Sort do a recursive sort over an outline
func (a ByText) Sort() {
	if len(a) > 0 {
		for _, item := range a {
			if len(item.Outline) > 0 {
				ol := ByText(item.Outline)
				ol.Sort()
			}
		}
		sort.Sort(ByText(a))
	}
}

// Len for ByTextCaseInsensitive sort of Outline
func (a ByTextCaseInsensitive) Len() int {
	return len(a)
}

// Swap for ByTextCaseInsensitive sort of Outline
func (a ByTextCaseInsensitive) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less for ByTextCaseInsensitive sort of Outline
func (a ByTextCaseInsensitive) Less(i, j int) bool {
	return strings.Compare(strings.ToUpper(a[i].Text), strings.ToUpper(a[j].Text)) == -1
}

// Sort do a recursive sort over an outline
func (a ByTextCaseInsensitive) Sort() {
	if len(a) > 0 {
		for _, item := range a {
			if len(item.Outline) > 0 {
				ol := ByTextCaseInsensitive(item.Outline)
				ol.Sort()
			}
		}
		sort.Sort(ByTextCaseInsensitive(a))
	}
}

// Len for ByTitle sort of Outline
func (a ByTitle) Len() int {
	return len(a)
}

// Swap for ByTitle sort of Outline
func (a ByTitle) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less for ByTitle sort of Outline
func (a ByTitle) Less(i, j int) bool {
	return strings.Compare(a[i].Title, a[j].Title) == -1
}

// Sort do a recursive sort over an outline
func (a ByTitle) Sort() {
	if len(a) > 0 {
		for _, item := range a {
			if len(item.Outline) > 0 {
				ol := ByTitle(item.Outline)
				ol.Sort()
			}
		}
		sort.Sort(ByTitle(a))
	}
}

// Len for ByTitleCaseInsensitive sort of Outline
func (a ByTitleCaseInsensitive) Len() int {
	return len(a)
}

// Swap for ByTitleCaseInsensitive sort of Outline
func (a ByTitleCaseInsensitive) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less for ByTitleCaseInsensitive sort of Outline
func (a ByTitleCaseInsensitive) Less(i, j int) bool {
	return strings.Compare(strings.ToUpper(a[i].Title), strings.ToUpper(a[j].Title)) == -1
}

// Sort do a recursive sort over an outline
func (a ByTitleCaseInsensitive) Sort() {
	if len(a) > 0 {
		for _, item := range a {
			if len(item.Outline) > 0 {
				ol := ByTitleCaseInsensitive(item.Outline)
				ol.Sort()
			}
		}
		sort.Sort(ByTitleCaseInsensitive(a))
	}
}

// Sort do a recursive ByText sort of outline elements starting at the OPML struct.
func (o *OPML) Sort() {
	if o.Body != nil && len(o.Body.Outline) > 0 {
		ol := ByText(o.Body.Outline)
		ol.Sort()
	}
}

// SortCaseInsensitive do a recursive ByTextCaseInsensitive sort of outline elements starting at the OPML struct.
func (o *OPML) SortCaseInsensitive() {
	if o.Body != nil && len(o.Body.Outline) > 0 {
		ol := ByTextCaseInsensitive(o.Body.Outline)
		ol.Sort()
	}
}

// SortTitle do a recusive ByTitle sort of outline elements starting at the OMPL struct
func (o *OPML) SortTitle() {
	if o.Body != nil && len(o.Body.Outline) > 0 {
		ol := ByTitle(o.Body.Outline)
		ol.Sort()
	}
}

// SortTitleCaseInsensitive do a recusive ByTitleCaseInsensitive sort of outline elements starting at the OMPL struct
func (o *OPML) SortTitleCaseInsensitive() {
	if o.Body != nil && len(o.Body.Outline) > 0 {
		ol := ByTitleCaseInsensitive(o.Body.Outline)
		ol.Sort()
	}
}

// Len for ByType sort of Outline
func (a ByType) Len() int {
	return len(a)
}

// Swap for ByType sort of Outline
func (a ByType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less for ByType sort of Outline
func (a ByType) Less(i, j int) bool {
	return strings.Compare(a[i].Type, a[j].Type) == -1
}

// Sort do a recursive sort over an outline
func (a ByType) Sort() {
	if len(a) > 0 {
		for _, item := range a {
			if len(item.Outline) > 0 {
				ol := ByType(item.Outline)
				ol.Sort()
			}
		}
		sort.Sort(ByType(a))
	}
}

// SortTypes do a recursive ByText sort of outline elements starting at the OPML struct.
func (o *OPML) SortTypes() {
	if o.Body != nil && len(o.Body.Outline) > 0 {
		ol := ByType(o.Body.Outline)
		ol.Sort()
	}
}

// Parse reads a []byte and returns a OMPL object and error
func Parse(src []byte) (*OPML, error) {
	o := New()
	err := xml.Unmarshal(src, &o)
	return o, err
}

// ReadFile reads a OPML file and returns a new OPML structure and error
func ReadFile(fname string) (*OPML, error) {
	o := New()
	err := o.ReadFile(fname)
	return o, err
}

// ReadFile reads an OPML file and populates the OPML object appropriately
func (o *OPML) ReadFile(s string) error {
	src, err := ioutil.ReadFile(s)
	if err != nil {
		return err
	}
	return xml.Unmarshal(src, &o)
}

// WriteFile writes the contents of a OPML struct to a file
func (o *OPML) WriteFile(s string, perm os.FileMode) error {
	if len(o.Body.Outline) == 0 {
		o.Body.Outline = append(o.Body.Outline, Outline{
			Text: "",
		})
	}
	b, _ := xml.Marshal(o)
	return ioutil.WriteFile(s, b, perm)
}

// Marshal takes an *OPML struct and returns a byte array of source and error
func Marshal(o *OPML) ([]byte, error) {
	return xml.Marshal(o)
}

// Unmarshal takes an OPML XML representation and populates an OPML struct
func Unmarshal(src []byte, o *OPML) error {
	return xml.Unmarshal(src, &o)
}
