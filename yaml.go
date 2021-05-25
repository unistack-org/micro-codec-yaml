// Package yaml provides a yaml codec
package yaml

import (
	"io"

	"github.com/ghodss/yaml"
	"github.com/unistack-org/micro/v3/codec"
	rutil "github.com/unistack-org/micro/v3/util/reflect"
)

type yamlCodec struct{}

const (
	flattenTag = "flatten"
)

func (c *yamlCodec) Marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case nil:
		return nil, nil
	case *codec.Frame:
		return m.Data, nil
	}

	if nv, nerr := rutil.StructFieldByTag(v, codec.DefaultTagName, flattenTag); nerr == nil {
		v = nv
	}
	return yaml.Marshal(v)
}

func (c *yamlCodec) Unmarshal(b []byte, v interface{}) error {
	if len(b) == 0 || v == nil {
		return nil
	}

	if m, ok := v.(*codec.Frame); ok {
		m.Data = b
		return nil
	}

	if nv, nerr := rutil.StructFieldByTag(v, codec.DefaultTagName, flattenTag); nerr == nil {
		v = nv
	}

	return yaml.Unmarshal(b, v)
}

func (c *yamlCodec) ReadHeader(conn io.Reader, m *codec.Message, t codec.MessageType) error {
	return nil
}

func (c *yamlCodec) ReadBody(conn io.Reader, v interface{}) error {
	if v == nil {
		return nil
	}

	buf, err := io.ReadAll(conn)
	if err != nil {
		return err
	} else if len(buf) == 0 {
		return nil
	}

	return c.Unmarshal(buf, v)
}

func (c *yamlCodec) Write(conn io.Writer, m *codec.Message, v interface{}) error {
	if v == nil {
		return nil
	}

	buf, err := c.Marshal(v)
	if err != nil {
		return err
	} else if len(buf) == 0 {
		return codec.ErrInvalidMessage
	}

	_, err = conn.Write(buf)
	return err
}

func (c *yamlCodec) String() string {
	return "yaml"
}

func NewCodec() codec.Codec {
	return &yamlCodec{}
}
