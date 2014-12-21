package message

import (
	"errors"

	"github.com/davecgh/go-xdr/xdr"
)

var ErrDecode = errors.New("message: XDR decode error")
var ErrProtoVer = errors.New("message: Invalid protocol version")
var ErrMagic = errors.New("message: Invalid magic")

// Decode the Header part of the xdr_encoded bytes,
// returns the header, remainder bytes and error
func DecodeHeader(xdr_enc []byte) (HeaderMsg, []byte, error) {

	var header HeaderMsg

	remainingBytes, err := xdr.Unmarshal(xdr_enc, &header)
	if err != nil {
		return header, remainingBytes, ErrDecode
	}

	if header.Version != PROTOCOL_VER {
		return header, remainingBytes, ErrProtoVer
	}

	if header.Magic != MSG_MAGIC && header.Magic != RELAY_MAGIC {
		return header, remainingBytes, ErrMagic
	}

	return header, remainingBytes, nil
}

// Decode the Positon part of the xdr_encoded bytes,
// returns the position, remainder bytes and error
func DecodePosition(xdr_enc []byte) (PositionMsg, []byte, error) {

	var position PositionMsg

	remainingBytes, err := xdr.Unmarshal(xdr_enc, &position)
	if err != nil {
		return position, remainingBytes, ErrDecode
	}

	return position, remainingBytes, nil
}

func BytesToString(bites []byte) string {
	for n, b := range bites {
		if b == 0 {
			return string(bites[:n])
		}
	}
	return string(bites[:])
}
