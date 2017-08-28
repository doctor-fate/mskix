package mskix

import (
	"fmt"
	"sync"

	"github.com/doctor-fate/mskix/device"
)

type Parser interface {
	Parse([]byte) (device.Data, error)
}

var parsers sync.Map

func Register(id device.ID, p Parser) {
	if p == nil {
		panic("mskix: Register p is nil")
	}
	if _, ok := parsers.LoadOrStore(id, p); ok {
		panic("mskix: Register called twice for device parser " + id)
	}
}

func Parsers() (devices []device.ID) {
	parsers.Range(func(k, v interface{}) bool {
		devices = append(devices, k.(device.ID))
		return true
	})
	return
}

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

func ParseWithID(id device.ID, data []byte) (device.Data, error) {
	v, ok := parsers.Load(id)
	if !ok {
		return device.Data{}, fmt.Errorf("mskix: ParseWithID unknown device parser %q (forgotten import?)", id)
	}
	p := v.(Parser)
	return p.Parse(data)
}
