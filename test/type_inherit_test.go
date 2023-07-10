package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}
func (m *MyManager) GetManagerName() string {
	return m.Name
}

type ViceDir interface {
	GetName() string
	GetViceDirName() string
}

type MyViceDir struct {
	Name string
}

func (m *MyViceDir) GetName() string {
	return m.Name
}
func (m *MyViceDir) GetViceDirName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Eric", GetName[Manager](&MyManager{Name: "Eric"}))
	assert.Equal(t, "Eko", GetName[ViceDir](&MyViceDir{Name: "Eko"}))
}
