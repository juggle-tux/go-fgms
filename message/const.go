package message

// Magic value for fgfs messages - currently FGFS
const MSG_MAGIC = 0x46474653 // "FGFS"

// Relay magix
const RELAY_MAGIC = 0x53464746 // GSGF

// Protocol Version is currently 1.1
const PROTOCOL_VER = 0x00010001

// Message Types
const (
	TYPE_CHAT  = 1 //= is this used ??
	TYPE_RESET = 6
	TYPE_POS   = 7
)

/* Message Sizes
XDR demands Id4 byte alignment, but some compilers use 8 byte alignment
so it's safe to let the overall size of a network message be a
multiple of 8!
*/
const (
	MAX_CALLSIGN_LEN   = 8
	MAX_CHAT_MSG_LEN   = 256
	MAX_MODEL_NAME_LEN = 96
	MAX_PROPERTY_LEN   = 52
)
