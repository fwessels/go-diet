package main

import (
	"testing"
)

func testEliminateAutogenerated(t *testing.T, code string, expected int) {

	file, _, _, _ := parse(fileSet, "", []byte(code), false)

	_, scanner := eliminateAutogenerated(file)

	if scanner.eligibleCount() != expected {
		t.Errorf("testEliminateAutogenerated(): \nexpected %d\ngot      %d", expected, scanner.eligibleCount())
	}

}

func TestNoAssignments(t *testing.T) {

	// test direct assignment to receiver
	testEliminateAutogenerated(t, valReceiverNoAssignments, 1)
}

func TestAssignment(t *testing.T) {

	// test direct assignment to receiver
	testEliminateAutogenerated(t, valReceiverAssignment, 0)
}

func TestAssignmentSelectorExpr(t *testing.T) {

	// test assignment to member of receiver
	testEliminateAutogenerated(t, valReceiverAssignmentSelectorExpr, 0)
}

func TestPtrNoAssignments(t *testing.T) {

	// already pointer assignment, so nothing to do
	testEliminateAutogenerated(t, ptrReceiverNoAssignments, 0)
}

const valReceiverNoAssignments = `package cmd

import (
	"fmt"
	"time"
)

type retryStorage struct {
	maxRetryAttempts int
	retryUnit        time.Duration
	retryCap         time.Duration
}

func (f retryStorage) ListDir(volume, path string) {
	fmt.Println(f)
}
`

const valReceiverAssignment = `package cmd

import (
	"time"
)

type retryStorage struct {
	maxRetryAttempts int
	retryUnit        time.Duration
	retryCap         time.Duration
}

func (f retryStorage) ListDir(volume, path string) {
	f = retryStorage{}
}
`

const valReceiverAssignmentSelectorExpr = `package cmd

import (
	"time"
)

type retryStorage struct {
	maxRetryAttempts int
	retryUnit        time.Duration
	retryCap         time.Duration
}

func (f retryStorage) ListDir(volume, path string) {
	f.maxRetryAttempts = 123
}
`

const ptrReceiverNoAssignments = `package cmd

import (
	"time"
)

type retryStorage struct {
	maxRetryAttempts int
	retryUnit        time.Duration
	retryCap         time.Duration
}

func (f *retryStorage) ListDir(volume, path string) {
	fmt.Println(f)
}
`
