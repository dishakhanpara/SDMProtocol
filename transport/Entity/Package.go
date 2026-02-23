package entity

import "net/netip"

type Package struct {
	Index   byte
	LAddr   netip.Addr
	RAddr   netip.Addr
	MesLen  byte
	Message string
}

func NewPackage(index byte, lAddr netip.Addr, rAddr netip.Addr, mesLen byte, message string) Package {
	return Package{
		index,
		lAddr,
		rAddr,
		mesLen,
		message,
	}
}
