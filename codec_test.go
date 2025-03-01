package yaml

import (
	"bytes"
	"testing"

	"go.unistack.org/micro/v4/codec"
)

func TestFrame(t *testing.T) {
	s := &codec.Frame{Data: []byte("test")}

	buf, err := NewCodec().Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf, []byte(`test`)) {
		t.Fatalf("bytes not equal %s != %s", buf, `test`)
	}
}

func TestFrameFlatten(t *testing.T) {
	s := &struct {
		One  string
		Name *codec.Frame `yaml:"name" codec:"flatten"`
	}{
		One:  "xx",
		Name: &codec.Frame{Data: []byte("test")},
	}

	buf, err := NewCodec(codec.Flatten(true)).Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf, []byte(`test`)) {
		t.Fatalf("bytes not equal %s != %s", buf, `test`)
	}
}

func TestNativeYamlTags(t *testing.T) {
	s := &struct {
		One string `yaml:"first"`
	}{
		One: "",
	}

	err := NewCodec().Unmarshal([]byte(`first: "val"`), s)
	if err != nil {
		t.Fatal(err)
	}
	if s.One != "val" {
		t.Fatalf("XXX %#+v\n", s)
	}
}
