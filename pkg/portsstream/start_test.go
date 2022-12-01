package portsstream

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestStream_Start(t *testing.T) {
	t.Run("Valid stream", func(t *testing.T) {
		f, err := os.Open("testdata/ports.json")
		if err != nil {
			t.Fatalf("Stream.Start() | error opening test data file: %v", err)
		}

		wantKey := "AEAJM"

		res := make(chan Result)
		s := &Stream{
			stream: res,
		}

		go func() {
			s.Start(context.Background(), f)
		}()

		for data := range res {
			if data.Data.Key != wantKey {
				t.Fatalf("Stream.Start() | want key %s, got %s", wantKey, data.Data.Key)
			}
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		f, err := os.Open("testdata/invalid.json")
		if err != nil {
			t.Fatalf("Stream.Start() | error opening test data file: %v", err)
		}

		res := make(chan Result)
		s := &Stream{
			stream: res,
		}

		go func() {
			s.Start(context.Background(), f)
		}()

		for data := range res {
			fmt.Println(data.Err)
			if data.Err == nil {
				t.Fatal("Stream.Start() | got err nil, want not nil")
			}
		}
	})
}
