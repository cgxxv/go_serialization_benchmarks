package goserbench

import (
	"time"
)

//easyjson:json
type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
	Map      map[string]interface{}
	MapPk    map[*string]interface{}
	MapSk    map[string]*int
	MapI     map[interface{}]interface{}
	MapPtr   *map[interface{}]interface{}
}

type NoTimeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

type NoTimeNoStringNoFloatA struct {
	Name     []byte
	BirthDay uint64
	Phone    []byte
	Siblings uint32
	Spouse   bool
	Money    uint64
}
