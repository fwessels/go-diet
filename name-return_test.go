package main

import (
	"testing"
)

func testShortIdent(t *testing.T, ident, expected string) {

	short := getShortIdent(ident)

	if short != expected {
		t.Errorf("testShortIdent(): \nexpected %s\ngot      %s", expected, short)
	}
}

func TestShortIdent(t *testing.T) {

	testShortIdent(t, "preSignValues", "psv")
	testShortIdent(t, "credentialHeader", "ch")
	testShortIdent(t, "signValues", "sv")
	testShortIdent(t, "APIErrorCode", "aec")
}
