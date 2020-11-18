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
	"sync"
)

type IpsecPeerParams struct {
	SaOutIdx uint32
	SaInIdx  uint32
	LocalEspSPI string
	LocalIntegKey string
	LocalEncrKey string
}

type Allocator interface {
	MechanismParams(srcIp, destIp string, saInIdx, saOutIdx uint32) IpsecPeerParams
	//SAIdx() string
	GenerateKey(uint8) string
	Restore(localIP, remoteIP string, saInIdx, saOutIdx uint32, espSPI, encrKey, integKey string)
}

type allocator struct {
	ipsecSAPeerMutex sync.RWMutex
	saIdxPool map[uint32]bool
	saIdxLast uint32
	ipsecSAPeers map[string]IpsecPeerParams
}

func NewAllocator() Allocator {
	return &allocator{
		ipsecSAPeerMutex: sync.RWMutex{},
		ipsecSAPeers: map[string]IpsecPeerParams{},
	}
}

func (a *allocator) MechanismParams(srcIp, destIp string, saInIdx, saOutIdx uint32) IpsecPeerParams {
	a.ipsecSAPeerMutex.Lock()
	defer a.ipsecSAPeerMutex.Unlock()
	// Hack because we can't tell whether we're src or dest
	srcdest := srcIp + "/" + destIp
	destsrc := destIp + "/" + srcIp
	ipsecParams,present := a.ipsecSAPeers[srcdest]
	if !present {
		ipsecParams, present = a.ipsecSAPeers[destsrc]
		if !present {
			allocSaOutIdx := a.saIdxLast
			if allocSaOutIdx < saOutIdx {
				allocSaOutIdx = saOutIdx
				a.saIdxLast = saOutIdx
			}
			a.saIdxLast += 1
			allocSaInIdx := a.saIdxLast
			if allocSaInIdx < saInIdx {
				allocSaInIdx = saInIdx
				a.saIdxLast = saInIdx
			}
			a.saIdxLast += 1
			ipsecParams = IpsecPeerParams{
				SaOutIdx:      allocSaOutIdx,
				SaInIdx:       allocSaInIdx,
				LocalEspSPI:   a.GenerateKey(8),
				LocalIntegKey: a.GenerateKey(20),
				LocalEncrKey:  a.GenerateKey(16),
			}
			// = len(a.ipsecSAPeers)
			a.ipsecSAPeers[srcdest] = ipsecParams
			a.ipsecSAPeers[destsrc] = ipsecParams
		}
	}
	return ipsecParams
}

// SA Index - Allocate a new SA Index
/*
func (a *allocator) SAIdx() string {
	for {
		idx := atomic.AddUint32(&a.saIdxLast, 1)
		if _, exists := a.saIdxPool.Load(idx); !exists {
			a.saIdxPool.Store(idx, true)
			return strconv.Itoa(int(idx))
		}
	}
}

 */

func (a *allocator) GenerateKey(size uint8) string {
	key := make([]byte, size)
	_, _ = rand.Read(key)

	return fmt.Sprintf("%x", key)
}

// Restore value of last Vni based on connections we have at the moment.
func (a *allocator) Restore(localIP, remoteIP string, saInIdx, saOutIdx uint32, espSPI, encrKey, integKey string) {
	a.ipsecSAPeerMutex.Lock()
	defer a.ipsecSAPeerMutex.Unlock()
	// Hack because we can't tell whether we're src or dest
	srcdest := localIP + "/" + remoteIP
	destsrc := remoteIP + "/" + localIP
	ipsecParams := IpsecPeerParams{
		SaOutIdx:      saOutIdx,
		SaInIdx:       saInIdx,
		LocalEspSPI:   espSPI,
		LocalIntegKey: integKey,
		LocalEncrKey:  encrKey,
	}
	a.ipsecSAPeers[srcdest] = ipsecParams
	a.ipsecSAPeers[destsrc] = ipsecParams
	if a.saIdxLast < saOutIdx {
		a.saIdxLast = saOutIdx
	}
	if a.saIdxLast < saInIdx {
		a.saIdxLast = saInIdx
	}
}
