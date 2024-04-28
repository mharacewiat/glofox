package storage

import "testing"

func TestSet(t *testing.T) {
	s := Storage[string, string]{}
	s.Set("Foo", "Bar")
	s.Set("Baz", "Qux")

	if !s.Has("Foo") {
		t.Errorf("Expected storage to contain \"Foo\", but it didn't")
	}
	if !s.Has("Baz") {
		t.Errorf("Expected storage to contain \"Baz\", but it didn't")
	}
}

func TestGet(t *testing.T) {
	s := Storage[string, string]{}
	s.Set("Foo", "Bar")

	if s.Get("Foo") != "Bar" {
		t.Errorf("Expected \"Bar\", but got %v", s.Get("Foo"))
	}
}
