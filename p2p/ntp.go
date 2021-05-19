// Copyright 2014 The MOAC-core Authors
// This file is part of the MOAC-core library.
//
// The MOAC-core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The MOAC-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the MOAC-core library. If not, see <http://www.gnu.org/licenses/>.

package p2p

import (
	"time"

	"github.com/beevik/ntp"
	"github.com/MOACChain/MoacLib/log"
)

var ntpServers = [4]string{
	"0.pool.ntp.org",
	"1.pool.ntp.org",
	"2.pool.ntp.org",
	"3.pool.ntp.org",
}

// getNtpTime gets current time from NTP servers and log it for future reference.
func getNtpTime() {
	for _, server := range ntpServers {
		options := ntp.QueryOptions{Timeout: 30 * time.Second}
		if response, err := ntp.QueryWithOptions(server, options); err == nil {
			ntptime := time.Now().Add(response.ClockOffset)
			log.Infof("[Current time] NTP: %v , LOCAL: %v", ntptime, time.Now())
			// one success per round is enough
			break
		} else {
			log.Infof("[Current time] err = %v", err)
		}
	}
}

// ntpCheck starts goroutine for logging NTP time.
func ntpCheck() {
	for {
		getNtpTime()
		time.Sleep(5 * time.Minute)
	}
}
