package ipsec

import (
	"sync"
	"testing"
)

func getIpSecParams(wg *sync.WaitGroup, t *testing.T, srcIp, destIp string, in, out uint32) {
	defer wg.Done()
	ia := NewAllocator()

	ipSecParams := ia.MechanismParams(srcIp, destIp, in, out)
	t.Logf("%s / %s -- %d, %d -- localSPI %s", srcIp, destIp,
		ipSecParams.SaInIdx, ipSecParams.SaOutIdx, ipSecParams.LocalEspSPI)
}

func TestIPSecAllocator(t *testing.T) {
	//g := NewWithT(t)
	var wg sync.WaitGroup
	wg.Add(1)
	go getIpSecParams(&wg, t,"1.1.1.1", "2.2.2.2", 1, 2)
	wg.Add(1)
	go getIpSecParams(&wg, t,"1.1.1.2", "2.2.2.3", 3, 4)
	wg.Add(1)
	go getIpSecParams(&wg, t,"1.1.1.3", "2.2.2.4", 3, 4)
	wg.Add(1)
	go getIpSecParams(&wg, t,"1.1.1.1", "2.2.2.2", 1, 2)
	wg.Wait()
}
