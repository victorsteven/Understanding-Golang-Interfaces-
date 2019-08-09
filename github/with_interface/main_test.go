package main

import (
	"errors"
	"reflect"
	"testing"
)

type FakeReleaseInfoer struct {
	Tag string
	Err error
}

func (f FakeReleaseInfoer) GetLatestReleaseTag(repo string) (string, error) {
	if f.Err != nil {
		return "", f.Err
	}

	return f.Tag, nil
}

func TestGetReleaseTagMessage(t *testing.T) {
	// f := FakeReleaseInfoer{
	// 	Tag: "v5.8.31",
	// 	Err: nil,
	// }
	// expectedMsg := "The latest release is v5.8.31"
	// msg, err := getReleaseTagMessage(f, "dev/null")
	// if err != nil {
	// 	t.Fatalf("Expected err to be nil but it was %s", err)
	// }
	// if expectedMsg != msg {
	// 	t.Fatalf("Expected %s but got %s", expectedMsg, msg)
	// }

	// Using table test
	cases := []struct {
		f           FakeReleaseInfoer
		repo        string
		expectedMsg string
		expectedErr error
	}{
		{
			f: FakeReleaseInfoer{
				Tag: "v5.8.31",
				Err: nil,
			},
			repo:        "doesnt/matter",
			expectedMsg: "The latest release is v5.8.31",
			expectedErr: nil,
		},
		{
			f: FakeReleaseInfoer{
				Tag: "v5.8.31",
				Err: errors.New("TCP timeout"),
			},
			repo:        "doesnt/foo",
			expectedMsg: "",
			expectedErr: errors.New("Error quering GitHub API: TCP timeout"),
		},
	}

	for _, c := range cases {
		msg, err := getReleaseTagMessage(c.f, c.repo)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}
		if c.expectedMsg != msg {
			t.Errorf("Expected %q but got %q", c.expectedMsg, msg)
		}
		t.Log("All test passed")
	}
}
