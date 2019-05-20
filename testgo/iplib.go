package main

import (
	"fmt"
	"net"
	"sort"

	"github.com/c-robinson/iplib"
)

func main() {
	// this calls net.ParseCIDR() under the hood, but returns an iplib.Net object
	_, ipna, err := iplib.ParseCIDR("192.168.1.0/22")
	if err != nil {
		// this will be an error from the net package
	}

	// NewNet() wants a net.IP and is waaaaaaaaaaaaaaaaay faster
	ipb := net.ParseIP("192.168.2.0")
	ipnb := iplib.NewNet(ipb, 22)

	// ...works for IPv6 too
	ipc := net.ParseIP("2001:db8::1")
	ipnc := iplib.NewNet(ipc, 64)

	fmt.Println(ipna.Count()) // 1022 -- good enough for ipv4, but...

	fmt.Println(ipnc.Count())  // 4294967295 -- ...sigh
	fmt.Println(ipnc.Count6()) // 18446744073709551616 -- yay Count6() !

	fmt.Println(iplib.CompareNets(ipna, ipnb)) // -1

	ipnlist := []iplib.Net{ipnb, ipna, ipnc}
	sort.Sort(iplib.ByNet(ipnlist)) // []iplib.Net{ ipna, ipnb, ipnc }

	elist := ipna.Enumerate(0, 0)
	fmt.Println(len(elist)) // 1022

	fmt.Println(ipna.ContainsNet(ipnb)) // true

	fmt.Println(ipna.NetworkAddress())   // 192.168.1.0
	fmt.Println(ipna.FirstAddress())     // 192.168.1.1
	fmt.Println(ipna.LastAddress())      // 192.168.3.254
	fmt.Println(ipna.BroadcastAddress()) // 192.168.3.255

	fmt.Println(ipnc.NetworkAddress())   // 2001:db8::1 -- meaningless in IPv6
	fmt.Println(ipnc.FirstAddress())     // 2001:db8::1
	fmt.Println(ipnc.LastAddress())      // 2001:db8::ffff:ffff:ffff:ffff
	fmt.Println(ipnc.BroadcastAddress()) // 2001:db8::ffff:ffff:ffff:ffff

	ipa1 := net.ParseIP("2001:db8::2")
	ipa1, err = ipna.PreviousIP(ipa1) //  net.IP{2001:db8::1}, nil
	ipa1, err = ipna.PreviousIP(ipa1) //  net.IP{}, ErrAddressAtEndOfRange
}
