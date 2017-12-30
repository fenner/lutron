// Copyright 2017 William C. Fenner. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lutron

import (
	"log"
	"strconv"
	"strings"
	"sync"
)

type OperatingMode int
type FanMode int
type CallStatus int
type DegreesF float64

// HVAC Controllers (
type Hvac struct {
	Component

	mu       sync.Mutex
	heatset  DegreesF
	coolset  DegreesF
	current	 DegreesF
	mode	 OperatingMode
	fan	 FanMode
	eco	 bool
	call	 CallStatus
}

// Footnote: "Allow for reported temperature values to be zero
// padded, with up to 3 digits before the decimal point, and
// 2 digits after."
// TODO: is ParseFloat OK with 001.10?
func parseCurrentTemp(s string) (DegreesF, error) {
	current, err := strconv.ParseFloat(s, 64)
	return DegreesF(current), err
}

func parseSetpoints(sheat string,scool string) (DegreesF, DegreesF, error) {
	heat, err := strconv.ParseFloat(sheat, 64)
	if err != nil {
		return 0, 0, err
	}
	cool, err := strconv.ParseFloat(scool, 64)
	return DegreesF(heat), DegreesF(cool), err
}

func parseMode(s string) (OperatingMode, error) {
	mode, err := strconv.Atoi(s)
	// TODO: range checking
	return OperatingMode(mode), err
}

func parseFan(s string) (FanMode, error) {
	fan, err := strconv.Atoi(s)
	// TODO: range checking
	return FanMode(fan), err
}

func parseEco(s string) (bool, error) {
	eco, err := strconv.Atoi(s)
	// TODO: range checking
	return eco == 2, err
}

func parseCall(s string) (CallStatus, error) {
	call, err := strconv.Atoi(s)
	// TODO: range checking
	return CallStatus(call), err
}

func (h *Hvac) handleEvent(event string) error {
	n := strings.Split(event, ",")
	action, err := strconv.Atoi(n[0])
	if err != nil {
		return err
	}
	switch action {
	case 1: // current temperature ºF
		temp, err := parseCurrentTemp(n[1])
		if err != nil {
			return err
		}
		h.handleCurrentTemp(temp)

	case 2: // heat and cool setpoints
		heat, cool, err := parseSetpoints(n[1],n[2])
		if err != nil {
			return err
		}
		h.handleSetpoints(heat, cool)

	case 3: // operating mode
		mode, err := parseMode(n[1])
		if err != nil {
			return err
		}
		h.handleMode(mode)

	case 4: // fan mode
		fan, err := parseFan(n[1])
		if err != nil {
			return err
		}
		h.handleFan(fan)

	case 5: // eco mode
		eco, err := parseEco(n[1])
		if err != nil {
			return err
		}
		h.handleEco(eco)

	case 14: // call status
		call, err := parseCall(n[1])
		if err != nil {
			return err
		}
		h.handleCall(call)

	case 15: // redundant current temperature ºC
		break

	case 16: // redundant setpoints ºC
		break

	default:
		log.Printf("HVAC %d ignoring %s", h.Id(), event)
	}
	return nil
}

func (h *Hvac) handleCurrentTemp(current DegreesF) {
}

func (h *Hvac) handleSetpoints(heat DegreesF, cool DegreesF) {
}

func (h *Hvac) handleMode(mode OperatingMode) {
}

func (h *Hvac) handleFan(fan FanMode) {
}

func (h *Hvac) handleEco(eco bool) {
}

func (h *Hvac) handleCall(call CallStatus) {
}


// this is all a big to-do...
func (h *Hvac) reconnect() {
	return
}
