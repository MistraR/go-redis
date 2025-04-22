package reply

import (
	"bytes"
	"go-redis/interface/resp"
	"strconv"
)

type ErrorReply interface {
	Error() string
	ToBytes() []byte
}

var (
	nullBulkReplyBytes = []byte("$-1")
	CRLF               = "\r\n"
)

type BulkReply struct {
	Arg []byte //"mistra" ="$6\r\nmistra\r\n"
}

func (b BulkReply) ToBytes() []byte {
	if len(b.Arg) == 0 {
		return nullBulkReplyBytes
	}
	return []byte("$" + strconv.Itoa(len(b.Arg)) + CRLF + string(b.Arg) + CRLF)
}

func MakeBulkReply(arg []byte) *BulkReply {
	return &BulkReply{Arg: arg}
}

type MultiBulkReply struct { //字符串数组回复
	Args [][]byte
}

func (m *MultiBulkReply) ToBytes() []byte {
	arglen := len(m.Args)
	var buf bytes.Buffer
	buf.WriteString("*" + strconv.Itoa(arglen) + CRLF)
	for _, arg := range m.Args {
		if arg == nil {
			buf.WriteString(string(nullBulkReplyBytes) + CRLF)
		} else {
			buf.WriteString("$" + strconv.Itoa(len(arg)) + CRLF + string(arg) + CRLF)
		}
	}
	return buf.Bytes()
}

func MakeMultiBulkReply(arg [][]byte) *MultiBulkReply {
	return &MultiBulkReply{Args: arg}
}

type StatusReply struct { //状态回复
	Status string
}

func (s *StatusReply) ToBytes() []byte {
	return []byte("+" + s.Status + CRLF)
}

func MakeStatusReply(status string) *StatusReply {
	return &StatusReply{Status: status}
}

type IntReply struct { //数字回复
	Code int64
}

func (s *IntReply) ToBytes() []byte {
	return []byte(":" + strconv.FormatInt(s.Code, 10) + CRLF)
}

func MakeIntReply(code int64) *IntReply {
	return &IntReply{Code: code}
}

type StandardErrReply struct { //通用错误回复
	Status string
}

func (s *StandardErrReply) Error() string {
	return s.Status
}

func (s *StandardErrReply) ToBytes() []byte {
	return []byte("-" + s.Status + CRLF)
}

func MakeStandardReply(status string) *StandardErrReply {
	return &StandardErrReply{Status: status}
}

/*
*
判断是否是异常回复
*/
func IsErrReply(reply resp.Reply) bool {
	return reply.ToBytes()[0] == '-'
}
