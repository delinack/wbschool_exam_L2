package main

import (
	"fmt"
	"testing"
	"time"
)

func TestOrChannel(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	t.Run("No channels", func(t *testing.T) {
		start := time.Now()
		select {
		case <-orChannel():
			t.Fatal("orChannel should not close with no input channels")
		case <-time.After(100 * time.Millisecond):
			fmt.Printf("No channels: done after %v\n", time.Since(start))
		}
	})

	t.Run("One channel", func(t *testing.T) {
		start := time.Now()
		<-orChannel(sig(1 * time.Second))
		fmt.Printf("One channel: done after %v\n", time.Since(start))
	})

	t.Run("Multiple channels", func(t *testing.T) {
		start := time.Now()
		<-orChannel(
			sig(1*time.Second),
			sig(5*time.Second),
		)
		fmt.Printf("Multiple channels: done after %v\n", time.Since(start))
	})
}
