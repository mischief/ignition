// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The vmware provider fetches a configuration from the VMware Guest Info
// interface.

package vmware

import (
	"time"

	"github.com/coreos/ignition/src/log"
	"github.com/coreos/ignition/src/providers"
)

type Creator struct{}

func (Creator) Create(logger *log.Logger) providers.Provider {
	return &provider{
		logger: logger,
	}
}

type provider struct {
	logger *log.Logger
}

func (p provider) ShouldRetry() bool {
	return false
}

func (p *provider) BackoffDuration() time.Duration {
	return 0
}
