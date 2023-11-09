package protocol

type uncPing struct {
	Time       uint8
	MAGIC      uint16
	ClientGUID uint8
}
