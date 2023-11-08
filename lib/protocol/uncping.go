package protocol

type uncping struct {
	Time       uint8
	MAGIC      uint16
	ClientGUID uint8
}
