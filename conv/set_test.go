package conv

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestToSet(t *testing.T) {
	assert := internal.NewAssert(t, "TestToSet")

	// To bool set.
	assert.Equal(map[bool]struct{}{true: {}, false: {}}, ToSet[bool]([]bool{true, true, false, false}))
	assert.Equal(map[bool]struct{}{true: {}}, ToSet[bool]([]int{1, 1, 1}))
	assert.Equal(map[bool]struct{}{false: {}}, ToSet[bool]([]string{"false", "0", "FALSE"}))

	_, err := ToSetE[bool]([]string{"foo", "1", "2", "2"})
	assert.IsNotNil(err)

	// To int set.
	assert.Equal(map[int]struct{}{0: {}, 1: {}}, ToSet[int]([]int{0, 1, 1}))
	assert.Equal(map[int]struct{}{0: {}, 1: {}}, ToSet[int]([]bool{false, true, true}))
	assert.Equal(map[int]struct{}{0: {}, 1: {}}, ToSet[int]([]string{"0", "1", "1"}))

	_, err = ToSetE[int]([]string{"foo", "1", "2"})
	assert.IsNotNil(err)

	// To float64 set.
	assert.Equal(map[float64]struct{}{0: {}, 1: {}}, ToSet[float64]([]int{0, 1, 1}))
	assert.Equal(map[float64]struct{}{0: {}, 1: {}}, ToSet[float64]([]bool{false, true, true}))
	assert.Equal(map[float64]struct{}{0: {}, 1: {}}, ToSet[float64]([]string{"0", "1", "1"}))

	_, err = ToSetE[float64]([]string{"foo", "1", "2"})
	assert.IsNotNil(err)

	// To string set.
	assert.Equal(map[string]struct{}{"CN": {}, "HK": {}, "AU": {}}, ToSet[string]([]string{"CN", "HK", "AU"}))
	assert.Equal(map[string]struct{}{"false": {}, "true": {}}, ToSet[string]([]bool{false, true, true}))
	assert.Equal(map[string]struct{}{"0": {}, "1": {}}, ToSet[string]([]int{0, 1, 1}))

	_, err = ToSetE[string]([]struct{}{{}})
	assert.IsNotNil(err)
}
