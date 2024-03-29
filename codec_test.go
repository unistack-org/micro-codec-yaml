package yaml

import (
	"bytes"
	"testing"

	"go.unistack.org/micro/v3/codec"
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

	buf, err := NewCodec().Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf, []byte(`test`)) {
		t.Fatalf("bytes not equal %s != %s", buf, `test`)
	}
}

func TestReadBody(t *testing.T) {
	s := &struct {
		Name string
	}{}
	c := NewCodec()
	b := bytes.NewReader(nil)
	err := c.ReadBody(b, s)
	if err != nil {
		t.Fatal(err)
	}
}
