package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode(msg string) ([]byte, error) {
	var length = int32(len(msg))
	var pkg = new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		fmt.Println("Error writing length to buffer: ", err)
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		fmt.Println("Error writing length to buffer: ", err)
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32

	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("Error reading length from buffer: ", err)
		return "", err
	}

	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(length)+4)
	_, err = reader.Read(pack)
	if err != nil {
		fmt.Println("Error reading from buffer: ", err)
		return "", err
	}

	return string(pack[4:]), nil
}
