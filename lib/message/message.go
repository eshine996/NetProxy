package message

import (
	"NetProxy/lib/tools"
	"bytes"
	"encoding/binary"
	"errors"
)

//定义消息类型常量
const (
	MSGVKEY  = "VKEY"
	MSGERROR = "ERRO"
	MSGHEART = "HEAR"
)

/*
	The message is formed as follows:
	+----+-----+---------+
	|type| len | content |
	+----+---------------+
	| 4  |  4  |   ...   |
	+----+---------------+
*/

//Message 定义一个消息类
type Message struct {
	MsgType string
	Content string
}

func NewMessage(msgType string, content string) Message {
	if len(msgType) > 4 && len(msgType) < 1 {
		panic(errors.New("infoType length  must < 4"))
	}
	return Message{
		MsgType: msgType,
		Content: content,
	}
}

func (m *Message) ToBytes() []byte {
	raw := bytes.NewBuffer([]byte{})

	_ = binary.Write(raw, binary.LittleEndian, []byte(m.MsgType))
	bContentLength := tools.Int2Bytes(len(m.Content))
	_ = binary.Write(raw, binary.LittleEndian, bContentLength)
	_ = binary.Write(raw, binary.LittleEndian, []byte(m.Content))
	return raw.Bytes()
}
