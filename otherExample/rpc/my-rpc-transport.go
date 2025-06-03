package main

import (
	"encoding/binary"
	"io"
	"net"
)

const headLen = 4

type Transport struct {
	conn net.Conn
}

func NewTransport(conn net.Conn) *Transport {
	return &Transport{conn: conn}
}

func (t *Transport) Send(data []byte) error {
	buf := make([]byte, headLen+len(data))
	binary.BigEndian.PutUint32(buf[:headLen], uint32(len(data)))
	copy(buf[headLen:], data)

	_, err := t.conn.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

func (t *Transport) Read() ([]byte, error) {
	header := make([]byte, headLen)
	_, err := io.ReadFull(t.conn, header)
	if err != nil {
		return nil, err
	}

	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(t.conn, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
