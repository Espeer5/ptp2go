/*******************************************************************************
*                                                                              *
*  pkg/reader/reader.go                                                        *
*                                                                              *
*  Defines the Ptp2FrameReader interface which reads in payloads from any ptp  *
*  data source, constructs it into a Ptp2Frame, and passes it along to the     *
*  backend.                                                                    *
*                                                                              *
*  Author:  Edward Speer                                                      *
*  Revised: 7/16/2025                                                          *
*                                                                              *
*******************************************************************************/

/*******************************************************************************
*  PACKAGE DECLARATION                                                         *
*******************************************************************************/

package reader

/*******************************************************************************
*  IMPORTS                                                                     *
*******************************************************************************/

import (
	"time"

	"github.com/Espeer5/ptp2go/pkg/ptp"
)

/*******************************************************************************
*  TYPE DEFINITIONS                                                            *
*******************************************************************************/

// Ptp2FrameReader is an interface for receiving PTPv2 frames from any source
// ReadFrame returns the next decoded from, local revieve timestamp, or err
type Ptp2FrameReader interface {
	// Read a frame from a source into a Ptp2Frame
	ReadFrame() (ptp.Ptp2Frame, time.Time, error)
}
