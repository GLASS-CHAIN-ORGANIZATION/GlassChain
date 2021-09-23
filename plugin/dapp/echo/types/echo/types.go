package echo

import (
	"reflect"

	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/types"
)

//          Action  
const (
	ActionPing = iota
	ActionPang
)

//          log  
const (
	TyLogPing = 100001
	TyLogPang = 100002
)

var (
	// EchoX       
	EchoX = "echo"
	//          Action    
	actionName = map[string]int32{
		"Ping": ActionPing,
		"Pang": ActionPang,
	}
	//        Log      
	logInfo = map[int64]*types.LogInfo{
		TyLogPing: {Ty: reflect.TypeOf(PingLog{}), Name: "PingLog"},
		TyLogPang: {Ty: reflect.TypeOf(PangLog{}), Name: "PangLog"},
	}
)
var elog = log.New("module", EchoX)

func init() {
	//              
	types.AllowUserExec = append(types.AllowUserExec, []byte(EchoX))
	types.RegFork(EchoX, InitFork)
	types.RegExec(EchoX, InitExecutor)
}

//InitFork ...
func InitFork(cfg *types.Chain33Config) {
	cfg.RegisterDappFork(EchoX, "Enable", 0)
}

//InitExecutor ...
func InitExecutor(cfg *types.Chain33Config) {
	types.RegistorExecutor(EchoX, NewType(cfg))
}

// Type         
type Type struct {
	types.ExecTypeBase
}

// NewType          
func NewType(cfg *types.Chain33Config) *Type {
	c := &Type{}
	c.SetChild(c)
	c.SetConfig(cfg)
	return c
}

// GetPayload            
func (b *Type) GetPayload() types.Message {
	return &EchoAction{}
}

// GetName         
func (b *Type) GetName() string {
	return EchoX
}

// GetTypeMap         action  ，      
func (b *Type) GetTypeMap() map[string]int32 {
	return actionName
}

// GetLogMap              ，  rpc      
func (b *Type) GetLogMap() map[int64]*types.LogInfo {
	return logInfo
}
