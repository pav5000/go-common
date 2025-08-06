package longdur

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_JSON(t *testing.T) {
	type Obj struct {
		LongDur Duration `json:"longdur"`
	}

	t.Run("Unmarshal regular", func(t *testing.T) {
		rawJSON := []byte(`
		{
			"longdur": "1y 2m 3d"
		}
		`)
		var v Obj

		err := json.Unmarshal(rawJSON, &v)

		require.NoError(t, err)
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
		rawJSON := []byte(`
		{
			"longdur": "\u0031y 2m 3d"
		}
		`)
		var v Obj

		err := json.Unmarshal(rawJSON, &v)

		require.NoError(t, err)
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

	t.Run("Unmarshal error", func(t *testing.T) {
		rawJSON := []byte(`
		{
			"longdur": "1p"
		}
		`)
		var v Obj

		err := json.Unmarshal(rawJSON, &v)

		require.Error(t, err)
	})

	t.Run("Marshal", func(t *testing.T) {
		rawJSON, err := json.Marshal(Obj{
			LongDur: Duration{
				Years:  1,
				Months: 2,
				Days:   3,
			},
		})

		require.NoError(t, err)
		assert.JSONEq(t, `{"longdur":"1y 2m 3d"}`, string(rawJSON))
	})

	t.Run("Marshal optional", func(t *testing.T) {
		type ObjOptional struct {
			LongDur *Duration `json:"longdur"`
		}

		rawJSON, err := json.Marshal(ObjOptional{
			LongDur: &Duration{
				Years:  1,
				Months: 2,
				Days:   3,
			},
		})

		require.NoError(t, err)
		assert.JSONEq(t, `{"longdur":"1y 2m 3d"}`, string(rawJSON))
	})
}
