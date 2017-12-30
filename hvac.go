// Copyright 2017 William C. Fenner. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lutron

import (
	"log"
	"strings"
)

type Hvac struct {
	Component
}

func (h *Hvac) handleEvent(event string) error {
	n := strings.Split(event, ",")
	log.Printf("HVAC got %v", n)
	log.Printf("HVAC %d ignoring %s", h.Id(), event)
	return nil
}

// this is all a big to-do...
func (h *Hvac) reconnect() {
	return
}
