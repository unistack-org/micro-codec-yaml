// Package yaml provides a yaml codec
package yaml // import "go.unistack.org/micro-codec-yaml/v3"

import (
	pb "go.unistack.org/micro-proto/v3/codec"
	"go.unistack.org/micro/v3/codec"
	rutil "go.unistack.org/micro/v3/util/reflect"
	yaml "gopkg.in/yaml.v3"
)

type yamlCodec struct {
	opts codec.Options
}

var _ codec.Codec = &yamlCodec{}

func (c *yamlCodec) Marshal(v interface{}, opts ...codec.Option) ([]byte, error) {
	if v == nil {
		return nil, nil
	}

	options := c.opts
	for _, o := range opts {
		o(&options)
	}

	if options.Flatten {
		if nv, nerr := rutil.StructFieldByTag(v, options.TagName, "flatten"); nerr == nil {
			v = nv
		}
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

	if options.Flatten {
		if nv, nerr := rutil.StructFieldByTag(v, options.TagName, "flatten"); nerr == nil {
			v = nv
		}
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

func (c *yamlCodec) String() string {
	return "yaml"
}

func NewCodec(opts ...codec.Option) *yamlCodec {
	return &yamlCodec{opts: codec.NewOptions(opts...)}
}
