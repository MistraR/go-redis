package database

import (
	"go-redis/interface/resp"
	"go-redis/lib/logger"
	"go-redis/resp/reply"
)

type EchoDatabase struct {
}

func NewEchoDatabase() *EchoDatabase {
	return &EchoDatabase{}
}

// 模拟redis内核，收到什么指令字符串 就回发什么字符串
func (e EchoDatabase) Exec(client resp.Connection, args [][]byte) resp.Reply {
	return reply.MakeMultiBulkReply(args)

}

func (e EchoDatabase) AfterClientClose(c resp.Connection) {
	logger.Info("EchoDatabase AfterClientClose")
}

func (e EchoDatabase) Close() {
	logger.Info("EchoDatabase Close")

}
