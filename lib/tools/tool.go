package tools

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

//Bytes2Int 将bytes转换为int
func Bytes2Int(b []byte) (int, error) {
	var str string
	for i := 0; i < len(b); i++ {
		str += strconv.Itoa(int(b[i]))
	}

	x, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return x, nil

}

//Int2Bytes 将int转为bytes
func Int2Bytes(i int) []byte {
	tmp := int32(i)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, &tmp)
	return bytesBuffer.Bytes()
}
