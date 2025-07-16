package reader

import (
    "time"

    "github.com/Espeer5/ptp2go/pkg/ptp"
)

// ptp2FrameReader is an interface for receiving PTPv2 frames from any source
// ReadFrame returns the next decoded from, local revieve timestamp, or err
type FrameReader interface {
    ReadFrame() (ptp.PtpFrame, time.Time, error)
}

