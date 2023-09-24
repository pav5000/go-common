package longdur

import (
	"errors"

	"gopkg.in/yaml.v3"
)

var _ yaml.Unmarshaler = &Duration{}

func (d *Duration) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.ScalarNode {
		return errors.New("cannot unmarshal non-scalar node into longdur.Duration")
	}

	parsed, err := Parse(node.Value)
	if err != nil {
		return err
	}

	*d = parsed
	return nil
}

var _ yaml.Marshaler = Duration{}

func (d Duration) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}
