package entity

type Package struct {
	Index   byte
	LAddr   string
	RAddr   string
	MesLen  byte
	Message string
}

func NewPackage(index byte, lAddr string, rAddr string, mesLen byte, message string) Package {
	return Package{
		index,
		lAddr,
		rAddr,
		mesLen,
		message,
	}
}
