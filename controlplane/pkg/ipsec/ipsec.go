// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipsec

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

type Allocator interface {
	SAIdx() string
	GenerateKey(uint8) string
	Restore(local_ip string, remote_ip string, vniId uint32)
}

type allocator struct {
	saIdxPool sync.Map
	saIdxLast uint32
}

func NewAllocator() Allocator {
	return &allocator{}
}

// Vni - Allocate a new VNI, odd if local_ip < remote_ip, even otherwise
func (a *allocator) SAIdx() string {
	for {
		idx := atomic.AddUint32(&a.saIdxLast, 1)
		if _, exists := a.saIdxPool.Load(idx); !exists {
			a.saIdxPool.Store(idx, true)
			return strconv.Itoa(int(idx))
		}
	}
}

func (a *allocator) GenerateKey(size uint8) string {
	key := make([]byte, size)
	_, _ = rand.Read(key)

	return fmt.Sprintf("%x", key)
}

// Restore value of last Vni based on connections we have at the moment.
func (a *allocator) Restore(localIP, remoteIP string, vniID uint32) {
	return
}
