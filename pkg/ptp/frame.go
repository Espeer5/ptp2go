/*******************************************************************************
*                                                                              *
*  pkg/ptp/frame.go                                                            *
*                                                                              *
*  Defines a Ptp2Frame, the basic unit of PTPv2 payload data. Provides         *
*  utilities for encoding and decoding data frames from big-endian byte        *
*  slices.                                                                     *
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
	"encoding/binary"
	"errors"
)

/*******************************************************************************
*  TYPE DEFINITIONS                                                            *
*******************************************************************************/

// MsgType identifies the four core PTPv2 message kinds
type MsgType uint8

const (
	Sync MsgType = iota
	FollowUp
	DelayReq
	DelayResp
)

// Minimal unit of PTPv2 payload data
type Ptp2Frame struct {
	Type      MsgType // Sync, FollowUp, DelayReq, DelayResp
	Sequence  uint16  // Sequence ID for matching msgs
	Timestamp uint64  // nanoseconds timestamp ptp t1 (FUp) or t3 (DelayResp)
}

/*******************************************************************************
*  FUNCTION DEFINITIONS                                                        *
*******************************************************************************/

// Encode serializes a Ptp2Frame into a fixed-length big-endian byte slice
func Encode(f Ptp2Frame) []byte {
	buf := make([]byte, 1+2+8)
	buf[0] = byte(f.Type)
	binary.BigEndian.PutUint16(buf[1:], f.Sequence)
	binary.BigEndian.PutUint64(buf[3:], f.Timestamp)
	return buf
}

// Decode parses a Ptp2Frame from a big-endian byte slice
func Decode(data []byte) (Ptp2Frame, error) {
	if len(data) < 11 {
		return Ptp2Frame{}, errors.New("ptp: frame too short")
	}
	return Ptp2Frame{
		Type:      MsgType(data[0]),
		Sequence:  binary.BigEndian.Uint16(data[1:3]),
		Timestamp: binary.BigEndian.Uint64(data[3:11]),
	}, nil
}
