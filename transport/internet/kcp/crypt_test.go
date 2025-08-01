package kcp_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/v4fly/v4ray-core/v0/common"
	. "github.com/v4fly/v4ray-core/v0/transport/internet/kcp"
)

func TestSimpleAuthenticator(t *testing.T) {
	cache := make([]byte, 512)

	payload := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	auth := NewSimpleAuthenticator()
	b := auth.Seal(cache[:0], nil, payload, nil)
	c, err := auth.Open(cache[:0], nil, b, nil)
	common.Must(err)
	if r := cmp.Diff(c, payload); r != "" {
		t.Error(r)
	}
}

func TestSimpleAuthenticator2(t *testing.T) {
	cache := make([]byte, 512)

	payload := []byte{'a', 'b'}

	auth := NewSimpleAuthenticator()
	b := auth.Seal(cache[:0], nil, payload, nil)
	c, err := auth.Open(cache[:0], nil, b, nil)
	common.Must(err)
	if r := cmp.Diff(c, payload); r != "" {
		t.Error(r)
	}
}
