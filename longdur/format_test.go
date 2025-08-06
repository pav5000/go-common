package longdur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String_AllZerosGiveEmptyString(t *testing.T) {
	dur := Duration{}

	res := dur.String()

	assert.Empty(t, res)
}

func Test_String_FilledOnlyYears_GiveStringWithOnlyYears(t *testing.T) {
	dur := Duration{
		Years: 1,
	}

	res := dur.String()

	assert.Equal(t, "1y", res)
}

func Test_String_FilledOnlyMonths_GiveStringWithOnlyMonths(t *testing.T) {
	dur := Duration{
		Months: 1,
	}

	res := dur.String()

	assert.Equal(t, "1m", res)
}

func Test_String_FilledOnlyDays_GiveStringWithOnlyDays(t *testing.T) {
	dur := Duration{
		Days: 1,
	}

	res := dur.String()

	assert.Equal(t, "1d", res)
}

func Test_String_FilledDifferentCombinations(t *testing.T) {
	t.Run("years days", func(t *testing.T) {
		dur := Duration{
			Years: 1,
			Days:  2,
		}

		res := dur.String()

		assert.Equal(t, "1y 2d", res)
	})

	t.Run("months days", func(t *testing.T) {
		dur := Duration{
			Months: 1,
			Days:   2,
		}

		res := dur.String()

		assert.Equal(t, "1m 2d", res)
	})

	t.Run("years months", func(t *testing.T) {
		dur := Duration{
			Years:  1,
			Months: 2,
		}

		res := dur.String()

		assert.Equal(t, "1y 2m", res)
	})

	t.Run("years months days", func(t *testing.T) {
		dur := Duration{
			Years:  1,
			Months: 2,
			Days:   3,
		}

		res := dur.String()

		assert.Equal(t, "1y 2m 3d", res)
	})
}
