// Package plantower implements functionality for taking readings from the
// Plantower family of sensors.
//
// Currently supported:
// - PMS5003
package plantower

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	startBytesSize = 2
	frameSize      = 32
)

// Reading contains the data returned in a single frame by the PMS5003.
type Reading struct {
	// "Standard Particle" aka lab environment
	pm1_lab   int
	pm2_5_lab int
	pm10_lab  int

	// Atmospheric measurements
	pm1_atm   int
	pm2_5_atm int
	pm10_atm  int

	// Particle Count per 0.1L of air
	pc_0_3 int
	pc_0_5 int
	pc_1   int
	pc_2_5 int
	pc_5   int
	pc_10  int
}

// ReadNext takes a reader providing a stream of bytes from the PMS5003 and
// returns a struct with the decoded values from a single reading.
//
// The stream does not necessarily have to begin with the start bytes.
func ReadNext(r io.Reader) (*Reading, error) {
	head := make([]byte, 1)
	for {
		_, err := r.Read(head)
		if err != nil {
			return nil, err
		}

		if head[0] == 0x42 {
			_, err := r.Read(head)
			if err != nil {
				return nil, err
			}

			if head[0] == 0x4d {
				break
			}
		}
	}

	frame := make([]byte, frameSize-startBytesSize)
	n, err := r.Read(frame)
	if err != nil {
		return nil, err
	}

	if n != frameSize-startBytesSize {
		return nil, errors.New("invalid read length")
	}

	pm1_lab := binary.BigEndian.Uint16(frame[2:4])
	pm2_5_lab := binary.BigEndian.Uint16(frame[4:6])
	pm10_lab := binary.BigEndian.Uint16(frame[6:8])

	pm1_atm := binary.BigEndian.Uint16(frame[8:10])
	pm2_5_atm := binary.BigEndian.Uint16(frame[10:12])
	pm10_atm := binary.BigEndian.Uint16(frame[12:14])

	pc_0_3 := binary.BigEndian.Uint16(frame[14:16])
	pc_0_5 := binary.BigEndian.Uint16(frame[16:18])
	pc_1 := binary.BigEndian.Uint16(frame[18:20])
	pc_2_5 := binary.BigEndian.Uint16(frame[20:22])
	pc_5 := binary.BigEndian.Uint16(frame[22:24])
	pc_10 := binary.BigEndian.Uint16(frame[24:26])

	// Reserved
	_ = frame[26:28]

	checksum := int(binary.BigEndian.Uint16(frame[28:30]))

	sum := 0x42 + 0x4d
	for _, v := range frame[:len(frame)-2] {
		sum += int(v)
	}

	if checksum != sum {
		return nil, fmt.Errorf("invalid checksum, expected: %d, got: %d", checksum, sum)
	}

	return &Reading{
		pm1_lab:   int(pm1_lab),
		pm2_5_lab: int(pm2_5_lab),
		pm10_lab:  int(pm10_lab),
		pm1_atm:   int(pm1_atm),
		pm2_5_atm: int(pm2_5_atm),
		pm10_atm:  int(pm10_atm),
		pc_0_3:    int(pc_0_3),
		pc_0_5:    int(pc_0_5),
		pc_1:      int(pc_1),
		pc_2_5:    int(pc_2_5),
		pc_5:      int(pc_5),
		pc_10:     int(pc_10),
	}, nil
}
