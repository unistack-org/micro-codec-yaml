// Package yaml provides a yaml codec
package yaml

import (
	"io"
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/unistack-org/micro/v3/codec"
)

type yamlCodec struct{}

func (c *yamlCodec) Marshal(b interface{}) ([]byte, error) {
	switch m := b.(type) {
	case nil:
		return nil, nil
	case *codec.Frame:
		return m.Data, nil
	}

	return yaml.Marshal(b)
}

func (c *yamlCodec) Unmarshal(b []byte, v interface{}) error {
	if len(b) == 0 {
		return nil
	}
	switch m := v.(type) {
	case nil:
		return nil
	case *codec.Frame:
		m.Data = b
		return nil
	}

	return yaml.Unmarshal(b, v)
}

func (c *yamlCodec) ReadHeader(conn io.Reader, m *codec.Message, t codec.MessageType) error {
	return nil
}

func (c *yamlCodec) ReadBody(conn io.Reader, b interface{}) error {
	switch m := b.(type) {
	case nil:
		return nil
	case *codec.Frame:
		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			return err
		} else if len(buf) == 0 {
			return nil
		}
		m.Data = buf
		return nil
	}

	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		return err
	} else if len(buf) == 0 {
		// not needed but similar changes in all codecs
		return nil
	}

	return yaml.Unmarshal(buf, b)
}

func (c *yamlCodec) Write(conn io.Writer, m *codec.Message, b interface{}) error {
	switch m := b.(type) {
	case nil:
		return nil
	case *codec.Frame:
		_, err := conn.Write(m.Data)
		return err
	}

	buf, err := yaml.Marshal(b)
	if err != nil {
		return err
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
