package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	run := func() {
		go func() {
			main()
		}()

		<-time.Tick(time.Second)
		os.Exit(0)
	}
	assert.NotPanics(t, run)
}
