package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Part 29 Reflect")
	fmt.Println("")

	fmt.Println("part 29.1 Penerapan reflect")
	var infc01A int = 2
	var infc01AReflect = reflect.ValueOf(infc01A)
	fmt.Printf("infc01A: %v\n", infc01A)
	fmt.Printf("infc01A Reflect: %v\n", infc01AReflect)
	fmt.Printf("infc01A Type: %v\n", reflect.TypeOf(infc01A))
	fmt.Printf("infc01AReflect Kind: %v\n", infc01AReflect.Kind())
	if infc01AReflect.Kind() == reflect.Int {
		fmt.Println("Tipe data is", reflect.Int)
	}
	var infc01B = []int{1, 4, 9, 16, 25}
	fmt.Printf("infc01B: %v\n", infc01B)
	var infc01BReflect = reflect.ValueOf(infc01B)
	fmt.Printf("infc01B Reflect: %v\n", infc01BReflect)
	fmt.Printf("infc01B Type: %v\n", infc01BReflect.Type())
	fmt.Printf("infc01B Kind: %v\n", infc01BReflect.Kind())
	if infc01BReflect.Kind() == reflect.Slice {
		fmt.Println("Tipe data is", reflect.Slice)
	}
	var infc01C = 25
	fmt.Printf("infc01C: %v\n", infc01C)
	var infc01CReflect = reflect.ValueOf(infc01C)
	fmt.Printf("infc01C Reflect: %v\n", infc01CReflect)
	fmt.Printf("infc01C Reflect int: %v\n", infc01CReflect.Int())
	fmt.Printf("infc01C Reflect interface: %v\n", infc01CReflect.Interface())
	var infc01D = infc01CReflect.Interface().(int)
	fmt.Printf("infc01D: %v\n", infc01D)
	var infc01DReflect = reflect.ValueOf(infc01D)
	fmt.Printf("infc01D Reflect: %v\n", infc01DReflect)
	fmt.Printf("infc01D Reflect Kind: %v\n", infc01DReflect.Kind())
	fmt.Println("")

	fmt.Println("Part 29.2 Pengaksesan properti variabel objek")
	infc02A := &student{"Noval", 7}
	fmt.Printf("infc02A: %v\n", infc02A)
	infc02A.getProperty()
	infc02B := student{"Adrie", 8}
	(&infc02B).getProperty()
	fmt.Println("")

	fmt.Println("Part 29.3 Pengaksesan method variabel object")
	infc03A := &student{"Deni", 6}
	fmt.Printf("infc03A: %v\n", infc03A)
	infc03AReflect := reflect.ValueOf(infc03A)
	fmt.Printf("infc03A Reflect: %v\n", infc03AReflect)
	var infc03AReflectMethod = infc03AReflect.MethodByName("setName")
	infc03AReflectMethod.Call([]reflect.Value{
		reflect.ValueOf("aaa"),
	})

	fmt.Printf("infc03A Reflect: %v\n", infc03A.name)
}

type student struct {
	name  string
	grade int
}

func (s *student) setName(name string) {
	s.name = name
}

func (s *student) getProperty() {
	reflectValue := reflect.ValueOf(s)
	fmt.Printf("reflectValue: %v\n", reflectValue)
	fmt.Printf("reflectValue kind: %v\n", reflectValue.Kind())
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	reflectType := reflectValue.Type()
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Printf("%v: %v\n", reflectType.Field(i).Name, reflectType.Field(i).Type)
	}
}
