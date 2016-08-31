// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugins

import (
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/soniah/gosnmp"
)

type NetworkPlugin struct{}

func (p NetworkPlugin) Fetch(snmp *gosnmp.GoSNMP) (map[string]float64, error) {
	oids := []string{
		// ".1.3.6.1.2.1.31.1.1.1.1", // ifName
		".1.3.6.1.2.1.31.1.1.1.6",  // ifHCInOctets
		".1.3.6.1.2.1.31.1.1.1.10", // ifHCOutOctets
	}
	log.Infof("[Net Plugin] Get SNMP data")
	result, err := snmp.Get(oids)
	if err != nil {
		return nil, fmt.Errorf("[Net Plugin] SNMP Error: %v", err)
	}
	return map[string]float64{
		"net-in":  float64(result.Variables[0].Value.(uint)),
		"net-out": float64(result.Variables[1].Value.(uint)),
	}, nil
}