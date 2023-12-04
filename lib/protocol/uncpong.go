package protocol

import (
	"encoding/binary"
	"io"
)

type uncPong struct {
	Time       uint8
	ClientGUID uint8
	MAGIC      uint16
	ServerID   uint8
}

func (p *uncPong) Read(loom io.Reader) error {

	err := binary.Read(loom, binary.BigEndian, &p.Time)
	if err != nil {
		return err
	}

	err = binary.Read(loom, binary.BigEndian, &p.ClientGUID)
	if err != nil {
		return err
	}

	err = binary.Read(loom, binary.BigEndian, &p.MAGIC)
	if err != nil {
		return err
	}

	return nil
}

func (p *uncPong) Write(loom io.Writer) error {

	err := binary.Write(loom, binary.BigEndian, &p.Time)
	if err != nil {
		return err
	}

	err = binary.Write(loom, binary.BigEndian, &p.ClientGUID)
	if err != nil {
		return err
	}

	err = binary.Write(loom, binary.BigEndian, &p.ServerID)
	if err != nil {
		return err
	}

	return nil
}
