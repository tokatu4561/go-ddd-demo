package fullname

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	firstName = "firstName"
	lastName  = "lastName"
)

func TestNewFullName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := NewFullName(firstName, lastName)
		if err != nil {
			log.Fatal(err)
		}

		want := &FullName{firstName: firstName, lastName: lastName}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(FullName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail firstName is empty", func(t *testing.T) {
		firstName := ""
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): firstName is required", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail firstName is invalid", func(t *testing.T) {
		firstName := "first$Name&"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): firstName has an invalid character. letter is only", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail lastName is empty", func(t *testing.T) {
		lastName := ""
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): lastName is required", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail lastName is invalid", func(t *testing.T) {
		lastName := "last$Name&"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): lastName has an invalid character. letter is only", firstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
