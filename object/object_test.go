package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("trues do not have same hash key")
	}

	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("falses do not have same hash key")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("true has same hash key as false")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integers with same content have different hashkeys")
	}

	two1 := &Integer{Value: 2}
	two2 := &Integer{Value: 2}

	if two1.HashKey() != two2.HashKey() {
		t.Errorf("integers with same content have different hashkeys")
	}

	if one1.HashKey() == two1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}
