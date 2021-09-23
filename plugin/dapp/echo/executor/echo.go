package executor

import (
	"github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	echotypes "github.com/33cn/plugin/plugin/dapp/echo/types/echo"
)

var (
	// KeyPrefixPing ping   
	KeyPrefixPing = "mavl-echo-ping:%s"
	// KeyPrefixPang pang   
	KeyPrefixPang = "mavl-echo-pang:%s"

	// KeyPrefixPingLocal local ping   
	KeyPrefixPingLocal = "LODB-echo-ping:%s"
	// KeyPrefixPangLocal local pang   
	KeyPrefixPangLocal = "LODB-echo-pang:%s"
)

// Init           ，         ，         0
func Init(name string, cfg *types.Chain33Config, sub []byte) {
	dapp.Register(cfg, echotypes.EchoX, newEcho, 0)
	InitExecType()
}

// InitExecType                    
func InitExecType() {
	ety := types.LoadExecutorType(echotypes.EchoX)
	ety.InitFuncList(types.ListMethod(&Echo{}))
}

// Echo        
type Echo struct {
	dapp.DriverBase
}

//             ，                      
func newEcho() dapp.Driver {
	c := &Echo{}
	c.SetChild(c)
	c.SetExecutorType(types.LoadExecutorType(echotypes.EchoX))
	return c
}

// GetDriverName           
func (h *Echo) GetDriverName() string {
	return echotypes.EchoX
}
