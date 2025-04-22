package reply

type PongReply struct {
}

var ponybytes = []byte("+PONG\r\n")

func (p PongReply) ToBytes() []byte {
	return ponybytes
}

func MakePongReply() *PongReply {
	return &PongReply{}
}

type OkReply struct {
}

var okBytes = []byte("+OK\r\n")

func (r *OkReply) ToBytes() []byte {
	return okBytes
}

var theOkReply = new(OkReply)

func MakeOkReply() *OkReply {
	return theOkReply
}

type NullBulkReply struct {
}

var nullBulkBytes = []byte("$-1\r\n") //空
func (n NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

var emptyMultiBulkBytes = []byte("*0\r\n") //空数组

type EmptyMultiBulkReply struct {
}

func (e *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

type NoReply struct {
}

var noBytes = []byte("") //空字符串

func (n NoReply) ToBytes() []byte {
	return noBytes
}
