/*******************************************************************************
*                                                                              *
*  pkg/ptp/frame.go                                                            *
*                                                                              *
*  Defines the Ptp2Handler struct, which runs the PTPv2 state machine.         *
*  Responds to incoming PTPv2 data payloads and sends out its own data to      *
*  implement the ptp function via an internal state machine.                   *
*                                                                              *
*  Author:  Edward Speer                                                       *
*  Revised: 7/17/2025                                                          *
*                                                                              *
*******************************************************************************/

/*******************************************************************************
*  PACKAGE DECLARATION                                                         *
*******************************************************************************/

package ptp

/*******************************************************************************
*  IMPORTS                                                                     *
*******************************************************************************/

import (
	"sync"
	"time"
)

/*******************************************************************************
*  TYPE DEFINITIONS                                                            *
*******************************************************************************/

// Ptp2Handler contains the state necessary to match and compute PTPv2 exchanges
type Ptp2Handler struct {
	lock sync.Mutex

	// The state for the most recent Sync+FollowUp sequence
	lastSync struct {
		seq      uint16
		t1recv   time.Time // local receive time of Sync
		t1p      uint64    // master timestamp from FollowUp
		have_fUp bool      // Has received matching FollowUp message
	}
}

// Constructor
func NewPtp2Handler() *Ptp2Handler {
	return &Ptp2Handler{}
}

// Handle ingests one Ptp2Frame along with the local receive/send time.
// For Sync/fUp it records t1, DelayRequest records t2, DelayResp computes
// offset and delay and calls report()
func (h *Ptp2Handler) Handle(frame Ptp2Frame, localTIme time.Time) error {
	// Protect the handler state via mutex until the end of the function call
	h.lock.Lock()
	defer h.lock.Unlock()

	// TODO: Implement PTPV2 state machine & compute offset/delay
	// This line is a placeholder to appease linter
	h.report(time.Duration(h.lastSync.t1p), time.Duration(h.lastSync.t1p))

	return nil
}

// report is a private method of the PTP2Handler which directs the offset and
// delay time to the configured output to allow usage of the timing data. For
// example this function could be used to direct the offset and delay
// information to an OCXO chip or other for clock synch
func (h *Ptp2Handler) report(offset, delay time.Duration) {
	println("OFFSET: ", offset.String(), ", DELAY: ", delay.String())

	// TODO: Provide output stream capability
}
