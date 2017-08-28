// Package mskix represents layer of abstraction between drawer and device parsers
package mskix

import (
	"fmt"
	"sync"

	"github.com/doctor-fate/mskix/device"
)

// Parser is the interface that must be implemented by a device parser.
type Parser interface {
	Parse([]byte) (device.Data, error)
}

var parsers sync.Map

// Register makes a parser available by specified id
// If Register is called twice with the same id or if p is nil,
// it panics.
func Register(id device.ID, p Parser) {
	if p == nil {
		panic("mskix: Register p is nil")
	}
	if _, ok := parsers.LoadOrStore(id, p); ok {
		panic("mskix: Register called twice for device parser " + id)
	}
}

// Parsers returns a list of the IDs of the registered parsers.
func Parsers() (devices []device.ID) {
	parsers.Range(func(k, v interface{}) bool {
		devices = append(devices, k.(device.ID))
		return true
	})
	return
}

// Parse parses input data and returns information suitable for drawing and error, if any
// Parse trying to find correct parser by calling Parse method on each registered parser.
func Parse(data []byte) (d device.Data, err error) {
	parsers.Range(func(k, v interface{}) bool {
		p := v.(Parser)
		if d, err = p.Parse(data); err == nil {
			return false
		}
		return true
	})
	if err != nil {
		err = fmt.Errorf("mskix: Parse no suitable parser found")
	}
	return
}

// ParseWithID behaves same as Parse but taking parser ID specified by user
func ParseWithID(id device.ID, data []byte) (device.Data, error) {
	v, ok := parsers.Load(id)
	if !ok {
		return device.Data{}, fmt.Errorf("mskix: ParseWithID unknown device parser %q (forgotten import?)", id)
	}
	p := v.(Parser)
	return p.Parse(data)
}
