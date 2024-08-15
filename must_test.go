package must_test

import (
	"errors"
	"testing"

	"github.com/uechoco/go-must"
)

func TestGet2(t *testing.T) {
	t.Parallel()

	sampleFunc := func(v string) (a string, err error) {
		if v == "get error" {
			return "", errors.New("unexpected arguments")
		}
		return "sample", nil
	}

	t.Run("not panic", func(t *testing.T) {
		got := must.Get2(sampleFunc("not error"))
		want := "sample"
		AssertEqual(t, want, got)
	})

	t.Run("panic", func(t *testing.T) {
		f := func() {
			_ = must.Get2(sampleFunc("get error"))
		}
		AssertPanic(t, f)
	})
}

func TestGet3(t *testing.T) {
	t.Parallel()

	sampleFunc := func(v string) (a string, b int, err error) {
		if v == "get error" {
			return "", 0, errors.New("unexpected arguments")
		}
		return "sample", 1, nil
	}

	t.Run("not panic", func(t *testing.T) {
		got1, got2 := must.Get3(sampleFunc("not error"))
		want1, want2 := "sample", 1
		AssertEqual(t, want1, got1)
		AssertEqual(t, want2, got2)
	})

	t.Run("panic", func(t *testing.T) {
		f := func() {
			_, _ = must.Get3(sampleFunc("get error"))
		}
		AssertPanic(t, f)
	})
}

func TestGet4(t *testing.T) {
	t.Parallel()

	sampleFunc := func(v string) (a string, b, c int, err error) {
		if v == "get error" {
			return "", 0, 0, errors.New("unexpected arguments")
		}
		return "sample", 1, 2, nil
	}

	t.Run("not panic", func(t *testing.T) {
		got1, got2, got3 := must.Get4(sampleFunc("not error"))
		want1, want2, want3 := "sample", 1, 2
		AssertEqual(t, want1, got1)
		AssertEqual(t, want2, got2)
		AssertEqual(t, want3, got3)
	})

	t.Run("panic", func(t *testing.T) {
		f := func() {
			_, _, _ = must.Get4(sampleFunc("get error"))
		}
		AssertPanic(t, f)
	})
}

func TestGet5(t *testing.T) {
	t.Parallel()

	sampleFunc := func(v string) (a string, b, c, d int, err error) {
		if v == "get error" {
			return "", 0, 0, 0, errors.New("unexpected arguments")
		}
		return "sample", 1, 2, 3, nil
	}

	t.Run("not panic", func(t *testing.T) {
		got1, got2, got3, got4 := must.Get5(sampleFunc("not error"))
		want1, want2, want3, want4 := "sample", 1, 2, 3
		AssertEqual(t, want1, got1)
		AssertEqual(t, want2, got2)
		AssertEqual(t, want3, got3)
		AssertEqual(t, want4, got4)
	})

	t.Run("panic", func(t *testing.T) {
		f := func() {
			_, _, _, _ = must.Get5(sampleFunc("get error"))
		}
		AssertPanic(t, f)
	})
}

func TestGet6(t *testing.T) {
	t.Parallel()

	sampleFunc := func(v string) (a string, b, c, d, e int, err error) {
		if v == "get error" {
			return "", 0, 0, 0, 0, errors.New("unexpected arguments")
		}
		return "sample", 1, 2, 3, 4, nil
	}

	t.Run("not panic", func(t *testing.T) {
		got1, got2, got3, got4, got5 := must.Get6(sampleFunc("not error"))
		want1, want2, want3, want4, want5 := "sample", 1, 2, 3, 4
		AssertEqual(t, want1, got1)
		AssertEqual(t, want2, got2)
		AssertEqual(t, want3, got3)
		AssertEqual(t, want4, got4)
		AssertEqual(t, want5, got5)
	})

	t.Run("panic", func(t *testing.T) {
		f := func() {
			_, _, _, _, _ = must.Get6(sampleFunc("get error"))
		}
		AssertPanic(t, f)
	})
}

func TestGet7(t *testing.T) {
	t.Parallel()

	sampleFunc := func(v string) (a string, b, c, d, e, f int, err error) {
		if v == "get error" {
			return "", 0, 0, 0, 0, 0, errors.New("unexpected arguments")
		}
		return "sample", 1, 2, 3, 4, 5, nil
	}

	t.Run("not panic", func(t *testing.T) {
		got1, got2, got3, got4, got5, got6 := must.Get7(sampleFunc("not error"))
		want1, want2, want3, want4, want5, want6 := "sample", 1, 2, 3, 4, 5
		AssertEqual(t, want1, got1)
		AssertEqual(t, want2, got2)
		AssertEqual(t, want3, got3)
		AssertEqual(t, want4, got4)
		AssertEqual(t, want5, got5)
		AssertEqual(t, want6, got6)
	})

	t.Run("panic", func(t *testing.T) {
		f := func() {
			_, _, _, _, _, _ = must.Get7(sampleFunc("get error"))
		}
		AssertPanic(t, f)
	})
}

// AssertEqual ensures that want and got values are equal.
func AssertEqual(t *testing.T, want, got interface{}) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

// AssertPanic ensures that f panics. If panic does not occur, it reports the failure.
func AssertPanic(t *testing.T, f func()) {
	t.Helper()
	didPanic := true
	defer func() {
		r := recover()
		if !didPanic {
			t.Errorf("AssertPanic failed: expected panic, but not panicked: %v", r)
		}
	}()

	f()
	didPanic = false
}
