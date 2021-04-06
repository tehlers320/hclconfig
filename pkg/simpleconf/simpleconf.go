package simpleconf

import (
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"fmt"
)

type Cfg struct {
	Module tfconfig.Module
	File string
}

func NewTFConf(terraformDir string) *Cfg { 
	Conf, err := LoadTFConfig(".")
	if err != nil {
		fmt.Errorf("error: %s", err)
	}
	Cfg := Cfg{
		Module: *Conf,
		File: terraformDir,
	}

	return &Cfg
}

func LoadTFConfig(terraformDir string) (*tfconfig.Module, error) {
	module, diags := tfconfig.LoadModule(terraformDir)

	if diags.HasErrors() {
		return &tfconfig.Module{}, diags.Err()
	}
	return module, diags.Err()
}

func (tfcfg *Cfg) FindTFDefault(tfvar string, Default string) string {
	var d interface{}
	if _, ok := tfcfg.Module.Variables[tfvar]; ok {
		d = tfcfg.Module.Variables[tfvar].Default
	} else {
		return Default
	}

	str := fmt.Sprintf("%v", d)
	// TODO check before is nil
	if str == "<nil>" {
		str = Default
	}

	return str
}
