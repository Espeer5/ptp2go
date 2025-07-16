/*******************************************************************************
*                                                                              *
*  pkg/ptp/frame_test.go                                                       *
*                                                                              *
*  Performs unit testing on the ptp2Frame object. Verifies that the object     *
*  can be constructed, populated, encoded, and decoded correctly.              *
*                                                                              *
*  Author:   Edward Speer                                                      *
*  Revised: 7/16/2025                                                          *
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
	"testing"

	"github.com/stretchr/testify/assert"
)

/*******************************************************************************
*  UNIT TESTS                                                                  *
*******************************************************************************/

// Tests that a user may construct a frame with data
func TestConstructFrame(t *testing.T) {
	// Construct a simple frame
	frame := ptp2Frame{
		Type:      Sync,
		Sequence:  0xBEEF,
		Timestamp: 0x1234567890ABCDEF,
	}

	// Verify the constructor succeeded
	assert.Equal(t, frame.Type, Sync, "Should init type")
	assert.Equal(t, frame.Sequence, uint16(0xBEEF), "Shouild init seqnum")
	assert.Equal(
		t,
		frame.Timestamp,
		uint64(0x1234567890ABCDEF),
		"Should init TS",
	)
}

// Tests that a user may encode and decode a frame
func TestEncodeDecodeFrame(t *testing.T) {
	// Construct a test frame
	origin := ptp2Frame{
		Type:      Sync,
		Sequence:  0xBEEF,
		Timestamp: 0x1234567890ABCDEF,
	}

	// Encoded the frame in big-endian byte slices
	encoded := Encode(origin)

	// Decode the frame without error and recover the correct values
	decoded, err := Decode(encoded)
	assert.NoError(t, err, "Decoding should not error on valid frame")
	assert.Equal(t, origin, decoded, "Decoding should recover valid frame")
}

// Ensures decode throws error on short frame
func TestDecodeShort(t *testing.T) {
	// Anything shorter than 11 bytes should error
	_, err := Decode([]byte{0x00, 0x01, 0x02})
	assert.Error(t, err, "Decode short frame should throw error")
}
