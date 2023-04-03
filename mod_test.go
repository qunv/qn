package qn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type ModStruct struct {
	reg
}

func (m *ModStruct) Lookup(p ProtocolType) Registry {
	//TODO implement me
	panic("implement me")
}

func NewModStruct() Api {
	return &ModStruct{}
}

func NewNotApiConstructor() (Api, string) {
	return &ModStruct{}, ""
}

func (m *ModStruct) Handle(r Request) Response {
	return nil
}

func TestMod(t *testing.T) {
	o := option{
		constructor: NewModStruct,
		rfType:      reflect.TypeOf(ModStruct{}),
	}
	assert.True(t, isApiConstructor(o.constructor))

	not := option{
		constructor: NewNotApiConstructor,
	}
	assert.False(t, isApiConstructor(not.constructor))
}
