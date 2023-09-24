package longdur

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_YAML(t *testing.T) {
	type Obj struct {
		LongDur Duration `yaml:"longdur"`
	}

	t.Run("Unmarshal regular", func(t *testing.T) {
		rawYAML := []byte(`longdur: 1y 2m 3d`)
		var v Obj

		err := yaml.Unmarshal(rawYAML, &v)

		assert.NoError(t, err)
		assert.Equal(t,
			Obj{
				LongDur: Duration{
					Years:  1,
					Months: 2,
					Days:   3,
				},
			},
			v,
		)
	})

	t.Run("Unmarshal quoted", func(t *testing.T) {
		rawYAML := []byte(`longdur: "1y 2m 3d"`)
		var v Obj

		err := yaml.Unmarshal(rawYAML, &v)

		assert.NoError(t, err)
		assert.Equal(t,
			Obj{
				LongDur: Duration{
					Years:  1,
					Months: 2,
					Days:   3,
				},
			},
			v,
		)
	})

	t.Run("Unmarshal escaped", func(t *testing.T) {
		rawYAML := []byte(`longdur: "\x31y 2m 3d"`)
		var v Obj

		err := yaml.Unmarshal(rawYAML, &v)

		assert.NoError(t, err)
		assert.Equal(t,
			Obj{
				LongDur: Duration{
					Years:  1,
					Months: 2,
					Days:   3,
				},
			},
			v,
		)
	})

	t.Run("Marshal", func(t *testing.T) {
		rawYAML, err := yaml.Marshal(Obj{
			LongDur: Duration{
				Years:  1,
				Months: 2,
				Days:   3,
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, "longdur: 1y 2m 3d\n", string(rawYAML))
	})
}
