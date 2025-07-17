/*******************************************************************************
*                                                                              *
*  pkg/ptp/handler_test.go                                                     *
*                                                                              *
*  Performs unit testing on the Ptp2Handler objects. Verifies that the         *
*  handler behaves as expected to implement the ptp function.                  *
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
	// Testing utilities
	"testing"

	"github.com/stretchr/testify/assert"

	// stdlib
	"time"
)

/*******************************************************************************
*  UNIT TESTS                                                                  *
*******************************************************************************/

// Tests that a user may instantiate a new Ptp2Handler and populate its fields
func TestConstructHandler(t *testing.T) {
	// Use the provided handler constructor
	handler := NewPtp2Handler()

	// Create a time for the latest Sync
	t1recv := time.Date(
		2025,        // year
		time.July,   // month
		17,          // day
		13,          // hour
		30,          // minute
		45,          // second
		123_456_789, // nanosecond
		time.UTC,    // location
	)

	// Populate all data in the handler
	handler.lastSync.seq = 0x01
	handler.lastSync.t1recv = t1recv
	handler.lastSync.t1p = 0xBEEF
	handler.lastSync.have_fUp = true

	// Attempt locking the handler
	handler.lock.Lock()
	defer handler.lock.Unlock()

	// Check that the data is correct
	assert.Equal(t, handler.lastSync.seq, uint16(0x01), "Incorrect Seq Num set")
	assert.Equal(t, handler.lastSync.t1recv, t1recv, "t1recv cannot init")
	assert.Equal(t, handler.lastSync.t1p, uint64(0xBEEF), "Cannot init t1")
	assert.Equal(t, handler.lastSync.have_fUp, true, "Cannot set have_fUp")
}

// Tests that alling Handler() does not throw an error
func TestHandleValid(t *testing.T) {
	// Construct a new handler
	handler := NewPtp2Handler()

	// Create a new Ptp2Frame to call handler.Handle() on
	frame := Ptp2Frame{
		Type:      Sync,
		Sequence:  0xBEEF,
		Timestamp: 0x1234567890ABCDEF,
	}

	// Create a local time to pass to handler.Handle()
	local_time := time.Date(
		2025,        // year
		time.July,   // month
		17,          // day
		13,          // hour
		30,          // minute
		45,          // second
		123_456_789, // nanosecond
		time.UTC,    // location
	)

	// Attempt to call Handle
	err := handler.Handle(frame, local_time)

	// Ensure no error was thrown
	assert.NoError(t, err)
}
