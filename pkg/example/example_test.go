package example

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Run("valid request", func(t *testing.T) {
		nickname := "Burglar"

		team := "Thorin and Company"

		req := Request{
			FirstName:     "Bilbo",
			LastName:      "Baggins",
			Email:         "bilbobaggins@gmail.com",
			Nickname:      &nickname,
			Team:          &team,
			Key:           bytes.Repeat([]byte("a"), 2048),
			Points:        999,
			Something:     0.01,
			SomethingElse: 0.001,
		}

		err := req.Validate()

		require.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		req := Request{
			FirstName: "bob",
			LastName:  "",
		}

		err := req.Validate()

		require.Error(t, err)
	})
}
