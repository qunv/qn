package controllers

import (
	"fmt"
	"reflect"
	"testing"
)

var r = reflect.TypeOf(GetSynonymApi{})

func Test_t1(t *testing.T) {
	fmt.Println("Type:", r.Name())
	fmt.Println("Kind:", r.Kind())

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		tag := field.Tag.Get("endpoint")

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
