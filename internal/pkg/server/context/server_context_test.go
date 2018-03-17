package context

import (
	"testing"
)

func TestNoAttribute(t *testing.T) {
	ctx := NewContext()

	attr, err := ctx.Attribute("attribute")

	if attr != nil {
		t.Error("Method shouldn't return an attribute")
	}

	if err != ErrNoAttribute {
		t.Error("There should be error")
	}
}

func TestSetGetAttribute(t *testing.T) {
	value := "Some text"
	name := "string"
	ctx := NewContext()

	ctx.SetAttribute(name, value)
	attr, err := ctx.Attribute(name)

	if attr != value {
		t.Errorf("SetAttribute(): expected: %s, actual: %s", value, attr)
	}

	if err == ErrNoAttribute {
		t.Error("There should be attribute")
	}
}

func TestSetGetAttributeStruct(t *testing.T) {
	type Something struct {
		name string
	}

	value := &Something{"String"}
	name := "struct"
	ctx := NewContext()

	ctx.SetAttribute(name, value)
	attr, err := ctx.Attribute(name)

	if attr != value {
		t.Errorf("SetAttribute(): expected: %s, actual: %s", value, attr)
	}

	if err == ErrNoAttribute {
		t.Error("There should be attribute")
	}
}

func TestDoubleSet(t *testing.T) {
	value1 := "String"
	value2 := "Other string"
	name := "string"

	ctx := NewContext()

	ctx.SetAttribute(name, value1)
	ctx.SetAttribute(name, value2)

	attr := ctx.Attribute(name)
	if attr != value2 {
		t.Errorf("SetAttribute(): expected: %s, actual: %s", value1, attr)
	}
}
