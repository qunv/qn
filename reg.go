package qn

import (
	"container/list"
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn/protocol"
	"reflect"
)

type Regs struct {
	regs map[protocol.ProtocolType]Registry
}

func (r Regs) Lookup(m protocol.ProtocolType) Registry {
	if v, ok := r.regs[m]; ok {
		return v
	}
	return nil
}

func Registers(regs ...Registry) Regs {
	rs := make(map[protocol.ProtocolType]Registry)
	for _, v := range regs {
		rs[v.GetMethod().GetProtocolType()] = v
	}
	return Regs{
		regs: rs,
	}
}

type reg struct {
	endpoint    string
	method      protocol.Method
	tags        []string
	middlewares []gin.HandlerFunc
}

func (r *reg) Tags(tags ...string) regBuilder {
	r.tags = tags
	return r
}

func (r *reg) MiddleWare(middlewares ...gin.HandlerFunc) regBuilder {
	r.middlewares = middlewares
	return r
}

func (r *reg) New() Registry {
	return r
}

type _withRegFunc func(reg *reg)

func (r *reg) GetEndpoint() string {
	return r.endpoint
}

func (r *reg) GetMethod() protocol.Method {
	return r.method
}

func (r *reg) GetTags() []string {
	return r.tags
}

var _typeOfReg = reflect.TypeOf(reg{})

func isReg(r reflect.Type) bool {
	return embedsType(r, _typeOfReg)
}

func embeds(i interface{}, e reflect.Type) bool {

	if i == nil {
		return false
	}

	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}

	return embedsType(t, e)

}

func embedsType(i reflect.Type, e reflect.Type) bool {
	types := list.New()
	types.PushBack(i)
	for types.Len() > 0 {
		t := types.Remove(types.Front()).(reflect.Type)

		if t == e {
			return true
		}

		if t.Kind() != reflect.Struct {
			continue
		}

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Anonymous {
				types.PushBack(f.Type)
			}
		}
	}
	return false
}
