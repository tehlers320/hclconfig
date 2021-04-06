package main

import (
	"fmt"
	"github.com/tehlers320/hclconf/pkg/simpleconf"
)



func main() {

	SimpleConf := simpleconf.NewTFConf(".")
	SimpleConf.FindTFDefault("x", "y")
	
}
