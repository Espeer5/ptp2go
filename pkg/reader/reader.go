package reader

import (
	"time"

	"github.com/Espeer5/ptp2go/pkg/ptp"
)

// Ptp2FrameReader is an interface for receiving PTPv2 frames from any source
// ReadFrame returns the next decoded from, local revieve timestamp, or err
type Ptp2FrameReader interface {
	ReadFrame() (ptp.Ptp2Frame, time.Time, error)
}
