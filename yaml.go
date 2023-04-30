// Package yaml provides a yaml codec
package yaml // import "go.unistack.org/micro-codec-yaml/v4"

import (
	"io"

	pb "go.unistack.org/micro-proto/v4/codec"
	"go.unistack.org/micro/v4/codec"
	rutil "go.unistack.org/micro/v4/util/reflect"
	"sigs.k8s.io/yaml"
)

type yamlCodec struct {
	opts codec.Options
}

var _ codec.Codec = &yamlCodec{}

const (
	flattenTag = "flatten"
)

func (c *yamlCodec) Marshal(v interface{}, opts ...codec.Option) ([]byte, error) {
	if v == nil {
		return nil, nil
	}

	options := c.opts
	for _, o := range opts {
		o(&options)
	}
	if nv, nerr := rutil.StructFieldByTag(v, options.TagName, flattenTag); nerr == nil {
		v = nv
	}

	switch m := v.(type) {
	case *codec.Frame:
		return m.Data, nil
	case *pb.Frame:
		return m.Data, nil
	}

	return yaml.Marshal(v)
}

func (c *yamlCodec) Unmarshal(b []byte, v interface{}, opts ...codec.Option) error {
	if len(b) == 0 || v == nil {
		return nil
	}

	options := c.opts
	for _, o := range opts {
		o(&options)
	}

	if nv, nerr := rutil.StructFieldByTag(v, options.TagName, flattenTag); nerr == nil {
		v = nv
	}

	switch m := v.(type) {
	case *codec.Frame:
		m.Data = b
		return nil
	case *pb.Frame:
		m.Data = b
		return nil
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

func NewCodec(opts ...codec.Option) *yamlCodec {
	return &yamlCodec{opts: codec.NewOptions(opts...)}
}
