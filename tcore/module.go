package tcore

import (
	tframework "tframework.com/rpc/tcore/interface"
)

//***************************************************
//author tim.huang
//2022/8/22
//
//
//***************************************************

//***********************    type    ****************************

//***********************    type_end    ****************************

//***********************    var    ****************************

//***********************    var_end    ****************************

//***********************    interface    ****************************

//***********************    interface_end    ****************************

// ***********************    struct    ****************************

type BaseModule struct {
	plugin int64
}

func (b *BaseModule) GetPlugin() int64 {
	return b.plugin
}

func (b *BaseModule) AddPlugin(plugin tframework.TServerPlugin) int64 {
	b.plugin = b.plugin | int64(plugin)
	return b.plugin
}

func (c *BaseModule) GetModuleName() (moduleName string) {
	moduleName = ""
	return
}

//***********************    struct_end    ****************************

func init() {
}
