package protocol

type uncpong struct {
	Time       uint8
	MAGIC      uint16
	ClientGUID uint8
}
