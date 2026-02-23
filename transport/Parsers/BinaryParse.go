package parsers

import (
	entity "SAPProto/transport/Entity"
	"errors"
)

func ParseToBinary(pack *entity.Package) ([]byte, error) {

	var message []byte = make([]byte, 1+
		(pack.LAddr.BitLen()/8)+
		(pack.RAddr.BitLen()/8)+
		int(pack.MesLen)+
		len(pack.Message))

	message[0] = pack.Index

	if pack.LAddr.Is4() {
		copy(message[1:], pack.LAddr.AsSlice())
	} else if pack.LAddr.Is6() {
		copy(message[1:], pack.LAddr.AsSlice())
	} else {
		return []byte{}, errors.New("Error LAaddr no Is4 no Is6")
	}

	return message, nil

}

func ParseToPackage(pack []byte) entity.Package {
	return entity.Package{}
}
