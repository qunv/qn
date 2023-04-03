package qn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIsReg(t *testing.T) {
	a := struct {
		reg `endpoint:"/v1/abc" method:"http_post" tag:"public,private"`
	}{}
	assert.True(t, isReg(reflect.TypeOf(a)))
}
