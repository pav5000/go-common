package longdur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse_EmptyStringGivesZeroDuration(t *testing.T) {
	dur, err := Parse("")

	assert.NoError(t, err)
	assert.Equal(t, Duration{0, 0, 0}, dur)
}

func Test_Parse_WHitespacesOnlyGiveZeroDuration(t *testing.T) {
	dur, err := Parse("  	")

	assert.NoError(t, err)
	assert.Equal(t, Duration{0, 0, 0}, dur)
}

func Test_Parse_YearsOnlyReturnFilledYearsField(t *testing.T) {
	dur, err := Parse("2y")

	assert.NoError(t, err)
	assert.Equal(t, Duration{
		Years: 2,
	}, dur)
}

func Test_Parse_MonthsOnlyReturnFilledMonthsField(t *testing.T) {
	dur, err := Parse("3m")

	assert.NoError(t, err)
	assert.Equal(t, Duration{
		Months: 3,
	}, dur)
}

func Test_Parse_DaysOnlyReturnFilledDaysField(t *testing.T) {
	dur, err := Parse("4d")

	assert.NoError(t, err)
	assert.Equal(t, Duration{
		Days: 4,
	}, dur)
}

func Test_Parse_AllGroupsFilled_ReturnAllFilledFields(t *testing.T) {
	dur, err := Parse("1d 2y 3m")

	assert.NoError(t, err)
	assert.Equal(t, Duration{
		Years:  2,
		Months: 3,
		Days:   1,
	}, dur)
}

func Test_Parse_AllGroupsAreMultiDigit_ReturnAllFilledFields(t *testing.T) {
	dur, err := Parse("123d 456y 789m")

	assert.NoError(t, err)
	assert.Equal(t, Duration{
		Years:  456,
		Months: 789,
		Days:   123,
	}, dur)
}

func Test_Parse_UnknownSuffixGivesError(t *testing.T) {
	_, err := Parse("1u")

	assert.ErrorContains(t, err, "unknown suffix 'u' in the group '1u'")
}

func Test_Parse_DuplicateGroupGivesError(t *testing.T) {
	t.Run("days", func(t *testing.T) {
		_, err := Parse("1d 2d")

		assert.ErrorContains(t, err, "duplicate group '2d'")
	})

	t.Run("months", func(t *testing.T) {
		_, err := Parse("1m 2m")

		assert.ErrorContains(t, err, "duplicate group '2m'")
	})

	t.Run("years", func(t *testing.T) {
		_, err := Parse("1y 2y")

		assert.ErrorContains(t, err, "duplicate group '2y'")
	})

	t.Run("all groups", func(t *testing.T) {
		_, err := Parse("1y 2m 3d 4d")

		assert.ErrorContains(t, err, "duplicate group '4d'")
	})
}
