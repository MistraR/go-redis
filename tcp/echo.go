package tcp

import (
	"bufio"
	"context"
	"go-redis/lib/logger"
	"go-redis/lib/sync/atomic"
	"go-redis/lib/sync/wait"
	"io"
	"net"
	"sync"
	"time"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func (e *EchoClient) Close() error {
	e.Waiting.WaitWithTimeout(10 * time.Second)
	_ = e.Conn.Close()
	return nil
}

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func MakeHandler() *EchoHandler {
	return &EchoHandler{}
}

func (handler *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if handler.closing.Get() {
		_ = conn.Close()
	}
	client := &EchoClient{
		Conn:    conn,
		Waiting: wait.Wait{},
	}
	handler.activeConn.Store(client, struct{}{}) //store并发方法 v=空接结构体 等价于hashset
	reader := bufio.NewReader(conn)
	for true {
		msg, err := reader.ReadString('\n') //一条完整消息以\n分割，拆包粘包问题
		if err != nil {
			if err == io.EOF { //客户端中断了连接
				logger.Info("Client Connect close")
				handler.activeConn.Delete(client)
			} else {
				logger.Warn(err)
			}
			return
		}
		client.Waiting.Add(1) //+1 表示自己正在处理业务，不要关闭当前连接
		b := []byte(msg)
		_, _ = conn.Write(b)  //回写数据
		client.Waiting.Done() //-1
	}
}

func (handler *EchoHandler) Close() error {
	logger.Info("handler shutting down")
	handler.closing.Set(true)
	handler.activeConn.Range(func(key, value any) bool { //遍历map
		client := key.(*EchoClient)
		_ = client.Conn.Close()
		return true //true 继续处理map中的下一个元素
	})
	return nil
}
