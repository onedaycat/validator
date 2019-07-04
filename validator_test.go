package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestISO8601DataTime(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v := New()
		v.ISO8601DataTime("data", "2019-01-02T10:11:12Z")
		v.ISO8601DataTime("data", "2019-01-02T10:11:12+07:00")
		require.True(t, v.IsValid())
	})

	t.Run("invalid", func(t *testing.T) {
		v := New()
		v.ISO8601DataTime("data", "2019-01-02T10:11:12")
		require.True(t, v.HasError())

		v = New()
		v.ISO8601DataTime("data", "")
		require.True(t, v.HasError())
	})
}

func TestEmail(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v := New()
		v.Email("data", "test.test@test.com")
		require.True(t, v.IsValid())

		v = New()
		v.Email("data", "test._test@test.com")
		require.True(t, v.IsValid())
	})

	t.Run("invalid", func(t *testing.T) {
		v := New()
		v.Email("data", "test@test@test.com")
		require.True(t, v.HasError())

		v = New()
		v.Email("data", "")
		require.True(t, v.HasError())
	})
}

func TestEqualString(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		v := New()
		v.EqualString("data", "กกกกกกกกกก", 10)
		require.True(t, v.IsValid())
	})

	t.Run("invalid", func(t *testing.T) {
		v := New()
		v.EqualString("data", "กกกกก", 10)
		require.True(t, v.HasError())

		v = New()
		v.EqualString("data", "", 10)
		require.True(t, v.HasError())
	})
}
