package ts

import (
	"time"

	"github.com/jakebluelee0/gomedia/core/media/av"
	"github.com/jakebluelee0/gomedia/core/media/format/ts/tsio"
)

// Stream struct
type Stream struct {
	av.CodecData

	demuxer *Demuxer
	muxer   *Muxer

	pid        uint16
	streamID   uint8
	streamType uint8

	tsw *tsio.TSWriter
	idx int

	iskeyframe bool
	pts, dts   time.Duration
	data       []byte
	datalen    int
}
