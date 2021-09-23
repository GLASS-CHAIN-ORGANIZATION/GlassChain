package rpc

import (
	"context"

	rpctypes "github.com/33cn/chain33/rpc/types"
	"github.com/33cn/chain33/types"
	echotypes "github.com/33cn/plugin/plugin/dapp/echo/types/echo"
)

// Jrpc        RPC      
type Jrpc struct {
	cli *channelClient
}

// RPC       
type channelClient struct {
	rpctypes.ChannelClient
}

// Init    rpc   
func Init(name string, s rpctypes.RPCServer) {
	cli := &channelClient{}
	//       ，     Jrpc，    grpc        
	cli.Init(name, s, &Jrpc{cli: cli}, nil)
}

// QueryPing                Query  ，      rpc Query            
//        ，           ，      
func (c *Jrpc) QueryPing(param *echotypes.Query, result *interface{}) error {
	if param == nil {
		return types.ErrInvalidParam
	}
	//                
	reply, err := c.cli.QueryPing(context.Background(), param)
	if err != nil {
		return err
	}
	*result = reply
	return nil
}

// QueryPing         
func (c *channelClient) QueryPing(ctx context.Context, queryParam *echotypes.Query) (types.Message, error) {
	return c.Query(echotypes.EchoX, "GetPing", queryParam)
}
