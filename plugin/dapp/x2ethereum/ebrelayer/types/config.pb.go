// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	config.proto
	relayer.proto

It has these top-level messages:
	SyncTxConfig
	Log
	RelayerConfig
	SyncTxReceiptConfig
	Deploy
*/
package types

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type SyncTxConfig struct {
	Chain33Host         string `protobuf:"bytes,1,opt,name=chain33host" json:"chain33host,omitempty"`
	PushHost            string `protobuf:"bytes,2,opt,name=pushHost" json:"pushHost,omitempty"`
	PushName            string `protobuf:"bytes,3,opt,name=pushName" json:"pushName,omitempty"`
	PushBind            string `protobuf:"bytes,4,opt,name=pushBind" json:"pushBind,omitempty"`
	MaturityDegree      int32  `protobuf:"varint,5,opt,name=maturityDegree" json:"maturityDegree,omitempty"`
	Dbdriver            string `protobuf:"bytes,6,opt,name=dbdriver" json:"dbdriver,omitempty"`
	DbPath              string `protobuf:"bytes,7,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache             int32  `protobuf:"varint,8,opt,name=dbCache" json:"dbCache,omitempty"`
	FetchHeightPeriodMs int64  `protobuf:"varint,9,opt,name=fetchHeightPeriodMs" json:"fetchHeightPeriodMs,omitempty"`
	StartSyncHeight     int64  `protobuf:"varint,10,opt,name=startSyncHeight" json:"startSyncHeight,omitempty"`
	StartSyncSequence   int64  `protobuf:"varint,11,opt,name=startSyncSequence" json:"startSyncSequence,omitempty"`
	StartSyncHash       string `protobuf:"bytes,12,opt,name=startSyncHash" json:"startSyncHash,omitempty"`
}

func (m *SyncTxConfig) Reset()         { *m = SyncTxConfig{} }
func (m *SyncTxConfig) String() string { return proto.CompactTextString(m) }
func (*SyncTxConfig) ProtoMessage()    {}

type Log struct {
	Loglevel        string `protobuf:"bytes,1,opt,name=loglevel" json:"loglevel,omitempty"`
	LogConsoleLevel string `protobuf:"bytes,2,opt,name=logConsoleLevel" json:"logConsoleLevel,omitempty"`
	LogFile         string `protobuf:"bytes,3,opt,name=logFile" json:"logFile,omitempty"`
	MaxFileSize     uint32 `protobuf:"varint,4,opt,name=maxFileSize" json:"maxFileSize,omitempty"`
	MaxBackups      uint32 `protobuf:"varint,5,opt,name=maxBackups" json:"maxBackups,omitempty"`
	MaxAge          uint32 `protobuf:"varint,6,opt,name=maxAge" json:"maxAge,omitempty"`
	LocalTime       bool   `protobuf:"varint,7,opt,name=localTime" json:"localTime,omitempty"`
	Compress        bool   `protobuf:"varint,8,opt,name=compress" json:"compress,omitempty"`
	CallerFile      bool   `protobuf:"varint,9,opt,name=callerFile" json:"callerFile,omitempty"`
	CallerFunction  bool   `protobuf:"varint,10,opt,name=callerFunction" json:"callerFunction,omitempty"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}

type RelayerConfig struct {
	Title               string        `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	SyncTxConfig        *SyncTxConfig `protobuf:"bytes,2,opt,name=syncTxConfig" json:"syncTxConfig,omitempty"`
	Log                 *Log          `protobuf:"bytes,3,opt,name=log" json:"log,omitempty"`
	JrpcBindAddr        string        `protobuf:"bytes,4,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	EthProvider         string        `protobuf:"bytes,5,opt,name=ethProvider" json:"ethProvider,omitempty"`
	BridgeRegistry      string        `protobuf:"bytes,6,opt,name=bridgeRegistry" json:"bridgeRegistry,omitempty"`
	Deploy              *Deploy       `protobuf:"bytes,7,opt,name=deploy" json:"deploy,omitempty"`
	EthMaturityDegree   int32         `protobuf:"varint,8,opt,name=ethMaturityDegree" json:"ethMaturityDegree,omitempty"`
	EthBlockFetchPeriod int32         `protobuf:"varint,9,opt,name=ethBlockFetchPeriod" json:"ethBlockFetchPeriod,omitempty"`
	EthProviderCli      string        `protobuf:"bytes,10,opt,name=ethProviderCli" json:"ethProviderCli,omitempty"`
}

func (m *RelayerConfig) Reset()         { *m = RelayerConfig{} }
func (m *RelayerConfig) String() string { return proto.CompactTextString(m) }
func (*RelayerConfig) ProtoMessage()    {}

func (m *RelayerConfig) GetSyncTxConfig() *SyncTxConfig {
	if m != nil {
		return m.SyncTxConfig
	}
	return nil
}

func (m *RelayerConfig) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *RelayerConfig) GetDeploy() *Deploy {
	if m != nil {
		return m.Deploy
	}
	return nil
}

type SyncTxReceiptConfig struct {
	Chain33Host       string `protobuf:"bytes,1,opt,name=chain33host" json:"chain33host,omitempty"`
	PushHost          string `protobuf:"bytes,2,opt,name=pushHost" json:"pushHost,omitempty"`
	PushName          string `protobuf:"bytes,3,opt,name=pushName" json:"pushName,omitempty"`
	PushBind          string `protobuf:"bytes,4,opt,name=pushBind" json:"pushBind,omitempty"`
	StartSyncHeight   int64  `protobuf:"varint,5,opt,name=startSyncHeight" json:"startSyncHeight,omitempty"`
	StartSyncSequence int64  `protobuf:"varint,6,opt,name=startSyncSequence" json:"startSyncSequence,omitempty"`
	StartSyncHash     string `protobuf:"bytes,7,opt,name=startSyncHash" json:"startSyncHash,omitempty"`
}

func (m *SyncTxReceiptConfig) Reset()         { *m = SyncTxReceiptConfig{} }
func (m *SyncTxReceiptConfig) String() string { return proto.CompactTextString(m) }
func (*SyncTxReceiptConfig) ProtoMessage()    {}

type Deploy struct {
	// 
	OperatorAddr string `protobuf:"bytes,1,opt,name=operatorAddr" json:"operatorAddr,omitempty"`
	//  
	DeployerPrivateKey string `protobuf:"bytes,2,opt,name=deployerPrivateKey" json:"deployerPrivateKey,omitempty"`
	// 
	ValidatorsAddr []string `protobuf:"bytes,3,rep,name=validatorsAddr" json:"validatorsAddr,omitempty"`
	// 
	InitPowers []int64 `protobuf:"varint,4,rep,name=initPowers" json:"initPowers,omitempty"`
}

func (m *Deploy) Reset()         { *m = Deploy{} }
func (m *Deploy) String() string { return proto.CompactTextString(m) }
func (*Deploy) ProtoMessage()    {}

func init() {
}
