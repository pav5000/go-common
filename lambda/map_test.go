package lambda

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapSlice_NilInput_ShouldGiveNilOutput(t *testing.T) {
	assert.Nil(t, MapSlice[int, int](nil, nil))
	assert.Nil(t, MapSlice[int, int](nil, func(i int) int { return 0 }))
}

func Test_MapSlice_RegularOutput_ShouldMapResults(t *testing.T) {
	assert.Equal(t,
		[]string{"1", "2", "3"},
		MapSlice([]int{1, 2, 3}, func(v int) string {
			return strconv.Itoa(v)
		}),
	)
}

func Test_SliceToMap_NilOrEmptyInput_ReturnsEmptyNotNilMap(t *testing.T) {
	{
		m := SliceToMap(nil, func(int) (int, bool) { return 0, true })
		assert.NotNil(t, m)
		assert.Empty(t, m)
	}

	{
		m := SliceToMap([]int{}, func(int) (int, bool) { return 0, true })
		assert.NotNil(t, m)
		assert.Empty(t, m)
	}
}

func Test_SliceToMap_RegularInput_MapsToCorrectOutput(t *testing.T) {
	assert.Equal(t,
		map[int]byte{
			1: 11,
			2: 12,
			3: 13,
		},
		SliceToMap([]int{1, 2, 3}, func(i int) (int, byte) { return i, byte(i + 10) }),
	)
}
