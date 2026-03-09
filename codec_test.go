package yaml

import (
	"bytes"
	"testing"

	yamlc "github.com/goccy/go-yaml"
	"go.unistack.org/micro/v4/codec"
	utime "go.unistack.org/micro/v4/util/time"
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

func TestDuration(t *testing.T) {
	s := &struct {
		One utime.Duration `yaml:"timeout"`
	}{
		One: 0,
	}
	err := NewCodec().Unmarshal([]byte(`timeout: "5s"`), s)
	if err != nil {
		t.Fatal(err)
	}
	if s.One != 5000000000 {
		t.Fatal("invalid duration parsed")
	}
}

func TestRawMessage(t *testing.T) {
	s := &struct {
		One *yamlc.RawMessage `yaml:"one"`
	}{}

	err := NewCodec().Unmarshal([]byte(`---
one:
  two:
    - test`), s)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("s %s", s.One)
}
