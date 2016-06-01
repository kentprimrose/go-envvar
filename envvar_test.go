package envvar

import (
	"os"
	"testing"
)

func TestFound(t *testing.T) {
	e := "FOUND"
	o := "original"
	err := os.Setenv(e, o)
	if err != nil {
		t.Fatalf("Error setting env: %s", err.Error())
	}

	v := os.Getenv(e)
	if v != o {
		t.Fatalf("Getenv: '%s' but '%s' expected", v, o)
	}

	d := "default"
	v, err = Getval(e, d)
	if err != nil {
		t.Fatalf("Getval (with default): Unexpected error - '%s'", err)
	}
	if v != o {
		t.Fatalf("Getval (with default): '%s' but '%s' expected", v, o)
	}

	v, err = Getval(e)
	if err != nil {
		t.Fatalf("Getval (without default): Unexpected error - '%s'", err)
	}
	if v != o {
		t.Fatalf("Getval (without default): '%s' but '%s' expected", v, o)
	}
}

func TestNotFound(t *testing.T) {
	e := "MISSING"
	os.Clearenv()

	v := os.Getenv(e)
	if v != "" {
		t.Fatalf("Getenv: '%s' not expected", v)
	}

	d := "default"
	v, err := Getval(e, d)
	if err != nil {
		t.Fatalf("Getval (with default): Unexpected error - '%s'", err)
	}
	if v != d {
		t.Fatalf("Getval (with default): '%s' but '%s' expected", v, d)
	}

	v, err = Getval(e)
	if err == nil {
		t.Fatal("Getval (without default): Expected error")
	}
}

func BenchmarkFoundWithDefault(b *testing.B) {
	e := "FOUND"
	o := "original"
	d := "default"
	err := os.Setenv(e, o)
	if err != nil {
		b.Fatalf("Error setting env: %s", err.Error())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Getval(e, d)
	}
}

func BenchmarkFoundWithoutDefault(b *testing.B) {
	e := "FOUND"
	o := "original"
	err := os.Setenv(e, o)
	if err != nil {
		b.Fatalf("Error setting env: %s", err.Error())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Getval(e)
	}
}

func BenchmarkNotFoundWithDefault(b *testing.B) {
	e := "FOUND"
	d := "default"
	os.Clearenv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Getval(e, d)
	}
}

func BenchmarkNotFoundWithoutDefault(b *testing.B) {
	e := "FOUND"
	os.Clearenv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Getval(e)
	}
}
