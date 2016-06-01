package goenv

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
	v, found := LookupEnv(e, d)
	if !found {
		t.Fatal("LookupEnv (with default): expected to find")
	}
	if v != o {
		t.Fatalf("LookupEnv (with default): '%s' but '%s' expected", v, o)
	}

	v, found = LookupEnv(e)
	if v != o {
		t.Fatalf("LookupEnv (without default): '%s' but '%s' expected", v, o)
	}
}

func TestNotFound(t *testing.T) {
	e := "MISSING"

	v := os.Getenv(e)
	if v != "" {
		t.Fatalf("Getenv: '%s' not expected", v)
	}

	d := "default"
	v, found := LookupEnv(e, d)
	if found {
		t.Fatal("LookupEnv (with default): Expected not to find")
	}
	if v != d {
		t.Fatalf("LookupEnv (with default): '%s' but '%s' expected", v, d)
	}

	v, found = LookupEnv(e)
	if found {
		t.Fatal("LookupEnv (without default): Expected not to find")
	}
}

func TestEmptyVal(t *testing.T) {
	e := "EMPTY"
	o := ""
	err := os.Setenv(e, o)
	if err != nil {
		t.Fatalf("Error setting env: %s", err.Error())
	}

	v := os.Getenv(e)
	if v != o {
		t.Fatalf("Getenv: '%s' but '%s' expected", v, o)
	}

	d := "default"
	v, found := LookupEnv(e, d)
	if !found {
		t.Fatal("LookupEnv (with default): Expected not to find")
	}
	if v != o {
		t.Fatalf("LookupEnv (with default): '%s' but '%s' expected", v, o)
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
		_, _ = LookupEnv(e, d)
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
		_, _ = LookupEnv(e)
	}
}

func BenchmarkNotFoundWithDefault(b *testing.B) {
	e := "MISSING"
	d := "default"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = LookupEnv(e, d)
	}
}

func BenchmarkNotFoundWithoutDefault(b *testing.B) {
	e := "MISSING"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = LookupEnv(e)
	}
}
