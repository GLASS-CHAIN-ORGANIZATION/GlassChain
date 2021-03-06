package executor

import (
	"bytes"
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/33cn/chain33/common"
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	ptypes "github.com/33cn/plugin/plugin/dapp/js/types"
	"github.com/33cn/plugin/plugin/dapp/js/types/jsproto"
	lru "github.com/hashicorp/golang-lru"
	"github.com/robertkrimen/otto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var driverName = ptypes.JsX
var basevm *otto.Otto
var codecache *lru.Cache

var isinit int64

//Init      
func Init(name string, cfg *types.Chain33Config, sub []byte) {
	if atomic.CompareAndSwapInt64(&isinit, 0, 1) {
		//   64 code cache
		var err error
		codecache, err = lru.New(512)
		if err != nil {
			panic(err)
		}
		drivers.Register(cfg, GetName(), newjs, 0)
	}
	InitExecType()
}

//InitExecType ...
func InitExecType() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&js{}))
	basevm = otto.New()
	_, err := basevm.Run(callcode)
	if err != nil {
		panic(err)
	}
	execaddressFunc(basevm)
	sha256Func(basevm)
}

type js struct {
	drivers.DriverBase
	prefix            []byte
	globalTableHandle sync.Map
	globalHanldeID    int64
}

func newjs() drivers.Driver {
	t := &js{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

//GetName     
func GetName() string {
	return newjs().GetName()
}

//GetDriverName        
func (u *js) GetDriverName() string {
	return driverName
}

func (u *js) ExecutorOrder() int64 {
	return drivers.ExecLocalSameTime
}

func (u *js) IsFriend(myexec, writekey []byte, othertx *types.Transaction) bool {
	if othertx == nil {
		return false
	}
	cfg := u.GetAPI().GetConfig()
	exec := cfg.GetParaExec(othertx.Execer)
	if exec == nil || len(bytes.TrimSpace(exec)) == 0 {
		return false
	}
	if string(exec) == ptypes.JsX {
		if bytes.HasPrefix(writekey, []byte("mavl-"+string(myexec))) {
			return true
		}
	}
	return false
}

func (u *js) callVM(prefix string, payload *jsproto.Call, tx *types.Transaction,
	index int, receiptData *types.ReceiptData) (*otto.Object, error) {
	if payload.Args != "" {
		newjson, err := rewriteJSON([]byte(payload.Args))
		if err != nil {
			return nil, err
		}
		payload.Args = string(newjson)
	} else {
		payload.Args = "{}"
	}
	loglist, err := jslogs(receiptData)
	if err != nil {
		return nil, err
	}
	vm, err := u.createVM(payload.Name, tx, index)
	if err != nil {
		return nil, err
	}
	vm.Set("loglist", loglist)
	if prefix == "init" {
		vm.Set("f", "init")
	} else {
		vm.Set("f", prefix+"_"+payload.Funcname)
	}
	vm.Set("args", payload.Args)
	callfunc := "callcode(context, f, args, loglist)"
	jsvalue, err := vm.Run(callfunc)
	//        ???         ???            ???           ???
	//               ???          ???           ???
	if u.GetExecutorAPI().IsErr() {
		return nil, status.New(codes.Aborted, "jsvm operation is abort").Err()
	}
	if err != nil {
		return nil, err
	}
	if prefix == "query" {
		s, err := jsvalue.ToString()
		if err != nil {
			return nil, err
		}
		return newObject(vm).setValue("result", s).object(), nil
	}
	if !jsvalue.IsObject() {
		return nil, ptypes.ErrJsReturnNotObject
	}
	return jsvalue.Object(), nil
}

type jslogInfo struct {
	Log    string `json:"log"`
	Ty     int32  `json:"ty"`
	Format string `json:"format"`
}

func jslogs(receiptData *types.ReceiptData) ([]string, error) {
	data := make([]string, 0)
	if receiptData == nil {
		return data, nil
	}
	for i := 0; i < len(receiptData.Logs); i++ {
		logitem := receiptData.Logs[i]
		//    json     ???          
		if logitem.Ty != ptypes.TyLogJs {
			continue
		}
		var jslog jsproto.JsLog
		err := types.Decode(logitem.Log, &jslog)
		if err != nil {
			return nil, err
		}
		item, err := json.Marshal(&jslogInfo{Log: jslog.Data, Ty: receiptData.Ty, Format: "json"})
		if err != nil {
			return nil, err
		}
		data = append(data, string(item))
	}
	return data, nil
}

func (u *js) getContext(tx *types.Transaction, index int64) *blockContext {
	var hash [32]byte
	if tx != nil {
		copy(hash[:], tx.Hash())
	}
	return &blockContext{
		Height:     u.GetHeight(),
		Name:       u.GetName(),
		Blocktime:  u.GetBlockTime(),
		Curname:    u.GetCurrentExecName(),
		DriverName: u.GetDriverName(),
		Difficulty: u.GetDifficulty(),
		TxHash:     common.ToHex(hash[:]),
		Index:      index,
		From:       tx.From(),
	}
}

func (u *js) statedbFunc(vm *otto.Otto, name string) {
	cfg := u.GetAPI().GetConfig()
	prefix, _ := calcAllPrefix(cfg, name)
	vm.Set("getstatedb", func(call otto.FunctionCall) otto.Value {
		key, err := call.Argument(0).ToString()
		if err != nil {
			return errReturn(vm, err)
		}
		hasprefix, err := call.Argument(1).ToBoolean()
		if err != nil {
			return errReturn(vm, err)
		}
		if !hasprefix {
			key = string(prefix) + key
		}
		v, err := u.getstatedb(key)
		if err != nil {
			return errReturn(vm, err)
		}
		return okReturn(vm, v)
	})
}

func (u *js) localdbFunc(vm *otto.Otto, name string) {
	cfg := u.GetAPI().GetConfig()
	_, prefix := calcAllPrefix(cfg, name)
	vm.Set("getlocaldb", func(call otto.FunctionCall) otto.Value {
		key, err := call.Argument(0).ToString()
		if err != nil {
			return errReturn(vm, err)
		}
		hasprefix, err := call.Argument(1).ToBoolean()
		if err != nil {
			return errReturn(vm, err)
		}
		if !hasprefix {
			key = string(prefix) + key
		}
		v, err := u.getlocaldb(key)
		if err != nil {
			return errReturn(vm, err)
		}
		return okReturn(vm, v)
	})
}

func (u *js) execnameFunc(vm *otto.Otto, name string) {
	cfg := u.GetAPI().GetConfig()
	vm.Set("execname", func(call otto.FunctionCall) otto.Value {
		return okReturn(vm, cfg.ExecName("user."+ptypes.JsX+"."+name))
	})
}

func (u *js) randnumFunc(vm *otto.Otto, name string) {
	vm.Set("randnum", func(call otto.FunctionCall) otto.Value {
		hash := u.GetLastHash()
		param := &types.ReqRandHash{
			ExecName: "ticket",
			BlockNum: 5,
			Hash:     hash,
		}
		randhash, err := u.GetExecutorAPI().GetRandNum(param)
		if err != nil {
			return errReturn(vm, err)
		}
		return okReturn(vm, common.ToHex(randhash))
	})
}

func (u *js) listdbFunc(vm *otto.Otto, name string) {
	//List(prefix, key []byte, count, direction int32) ([][]byte, error)
	cfg := u.GetAPI().GetConfig()
	_, plocal := calcAllPrefix(cfg, name)
	vm.Set("listdb", func(call otto.FunctionCall) otto.Value {
		prefix, err := call.Argument(0).ToString()
		if err != nil {
			return errReturn(vm, err)
		}
		key, err := call.Argument(1).ToString()
		if err != nil {
			return errReturn(vm, err)
		}
		count, err := call.Argument(2).ToInteger()
		if err != nil {
			return errReturn(vm, err)
		}
		direction, err := call.Argument(3).ToInteger()
		if err != nil {
			return errReturn(vm, err)
		}
		v, err := u.listdb(string(plocal)+prefix, key, int32(count), int32(direction))
		if err != nil {
			return errReturn(vm, err)
		}
		return listReturn(vm, v)
	})
}

func (u *js) createVM(name string, tx *types.Transaction, index int) (*otto.Otto, error) {
	data, err := json.Marshal(u.getContext(tx, int64(index)))
	if err != nil {
		return nil, err
	}
	var vm *otto.Otto
	if vmitem, ok := codecache.Get(name); ok {
		vm = vmitem.(*otto.Otto).Copy()
	} else {
		code, err := u.GetStateDB().Get(calcCodeKey(name))
		if err != nil {
			return nil, err
		}
		//cache       ???  cache     
		cachevm := basevm.Copy()
		cachevm.Run(code)
		codecache.Add(name, cachevm)
		vm = cachevm.Copy()
	}
	vm.Set("context", string(data))
	u.statedbFunc(vm, name)
	u.localdbFunc(vm, name)
	u.listdbFunc(vm, name)
	u.execnameFunc(vm, name)
	u.randnumFunc(vm, name)
	u.registerAccountFunc(vm)
	u.registerTableFunc(vm, name)
	return vm, nil
}

func errReturn(vm *otto.Otto, err error) otto.Value {
	return newObject(vm).setErr(err).value()
}

func okReturn(vm *otto.Otto, value string) otto.Value {
	return newObject(vm).setValue("value", value).value()
}

func listReturn(vm *otto.Otto, value []string) otto.Value {
	return newObject(vm).setValue("value", value).value()
}

func receiptReturn(vm *otto.Otto, receipt *types.Receipt) otto.Value {
	kvs := createKVObject(vm, receipt.KV)
	logs := createLogsObject(vm, receipt.Logs)
	return newObject(vm).setValue("kvs", kvs).setValue("logs", logs).value()
}

type object struct {
	vm  *otto.Otto
	obj *otto.Object
}

func newObject(vm *otto.Otto) *object {
	return newObjectString(vm, "({})")
}

func newObjectString(vm *otto.Otto, value string) *object {
	obj, err := vm.Object(value)
	if err != nil {
		panic(err)
	}
	return &object{vm: vm, obj: obj}
}

func (o *object) setErr(err error) *object {
	if err != nil {
		o.obj.Set("err", err.Error())
	}
	return o
}

func (o *object) setValue(key string, value interface{}) *object {
	o.obj.Set(key, value)
	return o
}

func (o *object) object() *otto.Object {
	return o.obj
}

func (o *object) value() otto.Value {
	v, err := otto.ToValue(o.obj)
	if err != nil {
		panic(err)
	}
	return v
}

// Allow               
func (u *js) Allow(tx *types.Transaction, index int) error {
	err := u.DriverBase.Allow(tx, index)
	if err == nil {
		return nil
	}
	//      :
	//  : user.jsvm.xxx     jsvm   
	//   : user.p.guodun.user.jsvm.xxx    jsvm   
	cfg := u.GetAPI().GetConfig()
	exec := cfg.GetParaExec(tx.Execer)
	if u.AllowIsUserDot2(exec) {
		return nil
	}
	return types.ErrNotAllow
}

func createKVObject(vm *otto.Otto, kvs []*types.KeyValue) []interface{} {
	data := make([]interface{}, len(kvs))
	for i := 0; i < len(kvs); i++ {
		item := make(map[string]interface{})
		item["key"] = string(kvs[i].Key)
		item["value"] = string(kvs[i].Value)
		item["prefix"] = true
		data[i] = item
	}
	return data
}

func createLogsObject(vm *otto.Otto, logs []*types.ReceiptLog) []interface{} {
	data := make([]interface{}, len(logs))
	for i := 0; i < len(logs); i++ {
		item := make(map[string]interface{})
		item["ty"] = logs[i].Ty
		item["log"] = string(logs[i].Log)
		item["format"] = "proto"
		data[i] = item
	}
	return data
}

func accountReturn(vm *otto.Otto, acc *types.Account) otto.Value {
	obj := newObject(vm)
	obj.setValue("currency", acc.Currency)
	obj.setValue("balance", acc.Balance)
	obj.setValue("frozen", acc.Frozen)
	obj.setValue("addr", acc.Addr)
	return obj.value()
}

func (u *js) getstatedb(key string) (value string, err error) {
	s, err := u.GetStateDB().Get([]byte(key))
	value = string(s)
	return value, err
}

func (u *js) getlocaldb(key string) (value string, err error) {
	s, err := u.GetLocalDB().Get([]byte(key))
	value = string(s)
	return value, err
}

func (u *js) listdb(prefix, key string, count, direction int32) (value []string, err error) {
	values, err := u.GetLocalDB().List([]byte(prefix), []byte(key), count, direction)
	for _, v := range values {
		value = append(value, string(v))
	}
	return value, err
}
