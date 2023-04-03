package qn

import (
	"reflect"
	"sync"
)

var tplPath *providedOption

func Init() {
	tplPath = new(providedOption)
}

func NewApi(fn func() (interface{}, reflect.Type)) {
	constructor, rt := fn()
	if !isApiConstructor(constructor) {
		panic("Is not an api constructor!")
	}

	if !isReg(rt) {
		panic("Missing dig constraints!")
	}

	o := option{
		constructor: constructor,
		rfType:      rt,
	}
	tplPath.append(o)
}

type option struct {
	constructor interface{}
	rfType      reflect.Type
}

type providedOption struct {
	mu      sync.Mutex
	options []option
}

func (p *providedOption) append(o option) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.options = append(p.options, o)
}

func (p *providedOption) path() []option {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.options
}

func isApiConstructor(v interface{}) bool {
	if reflect.TypeOf(v).Kind() == reflect.Func {
		ft := reflect.ValueOf(v).Type()
		if ft.NumOut() != 1 {
			return false
		}
		t := ft.Out(0)
		a := reflect.TypeOf(func(Api) {}).In(0)
		return t == a
	}
	return false
}
