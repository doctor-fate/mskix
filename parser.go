package mskix

import (
	"fmt"
	"sync"

	"github.com/doctor-fate/mskix/device"
)

type Parser interface {
	Parse(string) (device.Data, error)
}

var parsers sync.Map

func Register(id device.ID, p Parser) {
	if p == nil {
		panic("mskix: Register p is nil")
	}
	if _, ok := parsers.LoadOrStore(id, p); !ok {
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

func Parse(id device.ID, data string) (device.Data, error) {
	v, ok := parsers.Load(id)
	if !ok {
		return device.Data{}, fmt.Errorf("mskix: unknown device parser %q (forgotten import?)", id)
	}
	p := v.(Parser)
	return p.Parse(data)
}
