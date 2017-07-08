package klangsynthese

import (
	"testing"
	"time"

	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/synth"
	"github.com/200sc/klangsynthese/wav"
	"github.com/stretchr/testify/assert"
)

func TestFont1(t *testing.T) {
	f := NewFont()
	f.Filter(filter.Volume(.25))
	a, err := wav.NewController().Wave(synth.Sin(synth.A4, 2, 32))
	assert.Nil(t, err)
	fa := NewFontAudio(f, a)
	fa.Play()
	fa.Font.Filter(filter.Volume(.5))
	time.Sleep(2 * time.Second)
	fa.Play()
	fa.Font.Filter(filter.Volume(.75))
	time.Sleep(2 * time.Second)
	fa.Play()
	time.Sleep(2 * time.Second)

}
