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
	"hash/fnv"
	"sync"
)

/* IPSec param allocation via singleton pattern */
var (
	ipsecAllocMu = &sync.Mutex{}
	saIdxLast = uint32(1)
	saIdxPool = map[uint32]string{}
	ipsecSAPeers = map[string]*IpsecPeerParams{}
)

type IpsecPeerParams struct {
	SaOutIdx uint32
	SaInIdx  uint32
	LocalEspSPI string
	LocalIntegKey string
	LocalEncrKey string
}

type Allocator interface {
	MechanismParams(srcIp, destIp string, saInIdx, saOutIdx uint32) *IpsecPeerParams
	//SAIdx() string
	GenerateKey(uint8) string
	Restore(localIP, remoteIP string, saInIdx, saOutIdx uint32, espSPI, encrKey, integKey string)
}

type allocator struct {
	//ipsecSAPeerMutex sync.Mutex
	//saIdxPool map[uint32]bool
}

/*func init() {
	saIdxPool = map[uint32]bool{}
	saIdxLast = 1
	ipsecSAPeers = map[string]*IpsecPeerParams{}
}*/

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func ipPairTuple(ip1, ip2 string) string{
	return ip1 + "/" + ip2
}

func saIdxPoolInsert(idx uint32, srcDestTuple string) {
	saIdxPool[idx] = srcDestTuple
}

func getSaIdxPool(srcIp, destIp string, saInIdx, saOutIdx uint32) *IpsecPeerParams {
	ipPair1 := ipPairTuple(srcIp, destIp)
	ipPair2 := ipPairTuple(destIp, srcIp)
	inIpPair, inPresent := saIdxPool[saInIdx]
	if inPresent && (inIpPair == ipPair1 || inIpPair == ipPair2) {
		outIpPair, outPresent := saIdxPool[saOutIdx]
		if outPresent && (outIpPair == ipPair1 || outIpPair == ipPair2) {
			saPeer, peerPresent := ipsecSAPeers[ipPair1]
			if peerPresent {
				return saPeer
			}
		}
	}
	return nil
}

func NewAllocator() Allocator {
	return &allocator{
	}
}

func (a *allocator) checkInitSAIdx() {
	if saIdxLast == 0 {
		// start the allocation from 1
		saIdxLast = 1
	}
}

func (a *allocator) checkUpdateSAIdxLast(saIdx uint32) {
	if saIdxLast <= saIdx {
		saIdxLast = saIdx
		saIdxLast += 1
	}
}

func (a *allocator) MechanismParams(srcIp, destIp string, saInIdx, saOutIdx uint32) *IpsecPeerParams {
	ipsecAllocMu.Lock()
	defer ipsecAllocMu.Unlock()
	//a.checkInitSAIdx()
	//a.checkUpdateSAIdxLast(saOutIdx)
	//a.checkUpdateSAIdxLast(saInIdx)

	/*
	allocSaOutIdx := saOutIdx
	if allocSaOutIdx < saIdxLast {
		allocSaOutIdx = saIdxLast
		saIdxLast += 1
	} else {
		saIdxLast = allocSaOutIdx + 1
	}
	allocSaInIdx := saInIdx
	if allocSaInIdx < saIdxLast {
		allocSaInIdx = saIdxLast
		saIdxLast += 1
	} else {
		saIdxLast = allocSaInIdx + 1
	}
	*/
	// Hack because we can't tell whether we're src or dest
	srcdest := srcIp + "/" + destIp
	destsrc := destIp + "/" + srcIp
	/* Assume hash of srcdest & destsrc will be unique for the pool of connections */
	srcDestH := hash (srcdest)
	destSrcH := hash (destsrc)
	allocSaInIdx := srcDestH
	allocSaOutIdx := destSrcH
	if saInIdx == 0 && saOutIdx == 0 {
		allocSaInIdx = srcDestH
		allocSaOutIdx = destSrcH
	} else if saInIdx == destSrcH {
		allocSaInIdx = destSrcH
		allocSaOutIdx = srcDestH
	} else if saInIdx == srcDestH {
		allocSaInIdx = srcDestH
		allocSaOutIdx = destSrcH
	}

	ipsecParams,present := ipsecSAPeers[srcdest]
	if !present {
		ipsecParams, present = ipsecSAPeers[destsrc]
		if !present {
			ipsecParams = &IpsecPeerParams{
				SaOutIdx:      allocSaOutIdx,
				SaInIdx:       allocSaInIdx,
				LocalEspSPI:   a.GenerateKey(8),
				LocalIntegKey: a.GenerateKey(20),
				LocalEncrKey:  a.GenerateKey(16),
			}
			// = len(a.ipsecSAPeers)
			ipsecSAPeers[srcdest] = ipsecParams
			ipsecSAPeers[destsrc] = ipsecParams
		}
	}
	// ensure we adjust SA Index values if any passed in
	/*
	if saOutIdx != 0 || saInIdx != 0 {
		ipsecSAPeers[srcdest].SaOutIdx = allocSaOutIdx
		ipsecSAPeers[destsrc].SaInIdx = allocSaInIdx
	}
	*/
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
	ipsecAllocMu.Lock()
	defer ipsecAllocMu.Unlock()
	// Hack because we can't tell whether we're src or dest
	srcdest := localIP + "/" + remoteIP
	destsrc := remoteIP + "/" + localIP
	ipsecParams := &IpsecPeerParams{
		SaOutIdx:      saOutIdx,
		SaInIdx:       saInIdx,
		LocalEspSPI:   espSPI,
		LocalIntegKey: integKey,
		LocalEncrKey:  encrKey,
	}
	ipsecSAPeers[srcdest] = ipsecParams
	ipsecSAPeers[destsrc] = ipsecParams
	a.checkUpdateSAIdxLast(saOutIdx)
	a.checkUpdateSAIdxLast(saInIdx)
}
