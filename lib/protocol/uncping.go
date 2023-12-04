package protocol

import (
	"encoding/binary"
	"io"
)

type uncPing struct {
	Time       uint8
	MAGIC      uint16
	ClientGUID uint8
}

func (p *uncPing) Read(loom io.Reader) error {
	err := binary.Read(loom, binary.BigEndian, &p.Time)
	if err != nil {
		return err
	}
	err = binary.Read(loom, binary.BigEndian, &p.MAGIC)
	if err != nil {
		return err
	}
	err = binary.Read(loom, binary.BigEndian, &p.ClientGUID)
	if err != nil {
		return err
	}
	return nil
}

func (p *uncPing) Write(loom io.Writer) error {
	err := binary.Write(loom, binary.BigEndian, &p.Time)
	if err != nil {
		return err
	}
	err = binary.Write(loom, binary.BigEndian, &p.MAGIC)
	if err != nil {
		return err
	}
	err = binary.Write(loom, binary.BigEndian, &p.ClientGUID)
	if err != nil {
		return err
	}
	return nil
}
