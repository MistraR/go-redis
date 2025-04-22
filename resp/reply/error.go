package reply

type UnknownErrReply struct {
}

var unknownErrBytes = []byte("-Err unknown\r\n")

func (u UnknownErrReply) Error() string {
	return "Err unknown"
}

func (u UnknownErrReply) ToBytes() []byte {
	return unknownErrBytes
}

type ArgNumErrReply struct {
	Cmd string
}

func (a *ArgNumErrReply) Error() string {
	return ("-Err wrong number of arguments for '" + a.Cmd + "' command\r\n")
}

func (a *ArgNumErrReply) ToBytes() []byte {
	return []byte("-Err wrong number of arguments for '" + a.Cmd + "' command\r\n")
}

func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{Cmd: cmd}
}

type SyntaxErrReply struct {
}

var syntaxErrBytes = []byte("-Err syntax error\r\n")

var theSyntaxErrReply = &SyntaxErrReply{}

func MakeSyntaxErrReply(cmd string) *SyntaxErrReply {
	return theSyntaxErrReply
}

func (a *SyntaxErrReply) Error() string {
	return ("Err syntax error")
}

func (a *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

type WrongTypeErrReply struct {
}

func (w *WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key holding the wrong kind of value"
}

func (w *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")

type ProtocolErrReply struct {
	Msg string
}

func (p *ProtocolErrReply) Error() string {
	return "ERR Protocol error:" + p.Msg
}

func (p *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error:'" + p.Msg + "'\r\n")
}
