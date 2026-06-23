package v2

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollect(t *testing.T) {
	t.Run("returns all items from a single page", func(t *testing.T) {
		var calls int
		items, err := Collect(func(page int) ([]int, APIResponseMeta, error) {
			calls++
			return []int{1, 2, 3}, APIResponseMeta{CurrentPage: 1, LastPage: 1}, nil
		})

		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, items)
		assert.Equal(t, 1, calls, "fetch should be called once for a single page")
	})

	t.Run("concatenates items across pages in order", func(t *testing.T) {
		pages := [][]int{{1, 2}, {3, 4}, {5}}
		var requested []int
		items, err := Collect(func(page int) ([]int, APIResponseMeta, error) {
			requested = append(requested, page)
			return pages[page-1], APIResponseMeta{CurrentPage: page, LastPage: len(pages)}, nil
		})

		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
		assert.Equal(t, []int{1, 2, 3}, requested, "fetch should walk pages 1..LastPage in order")
	})

	t.Run("aborts atomically when a page errors", func(t *testing.T) {
		boom := errors.New("boom")
		var calls int
		items, err := Collect(func(page int) ([]int, APIResponseMeta, error) {
			calls++
			if page == 2 {
				return nil, APIResponseMeta{}, boom
			}
			return []int{page}, APIResponseMeta{CurrentPage: page, LastPage: 3}, nil
		})

		require.ErrorIs(t, err, boom)
		assert.Nil(t, items, "a mid-pagination error must yield no partial results")
		assert.Equal(t, 2, calls, "fetch should stop at the failing page")
	})

	t.Run("treats absent pagination metadata as a single page", func(t *testing.T) {
		var calls int
		items, err := Collect(func(page int) ([]int, APIResponseMeta, error) {
			calls++
			return []int{7, 8}, APIResponseMeta{}, nil
		})

		require.NoError(t, err)
		assert.Equal(t, []int{7, 8}, items)
		assert.Equal(t, 1, calls)
	})

	t.Run("stops on an empty page even if more are advertised", func(t *testing.T) {
		var calls int
		items, err := Collect(func(page int) ([]int, APIResponseMeta, error) {
			calls++
			if page == 1 {
				return []int{1}, APIResponseMeta{CurrentPage: 1, LastPage: 10}, nil
			}
			return nil, APIResponseMeta{CurrentPage: page, LastPage: 10}, nil
		})

		require.NoError(t, err)
		assert.Equal(t, []int{1}, items)
		assert.Equal(t, 2, calls, "fetch should stop once a page comes back empty")
	})
}
