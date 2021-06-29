package main

import (
	"context"
	"flag"
	"github.com/mpon/terraform-provider-twilio/twilio"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()
	opts := &plugin.ServeOpts{ProviderFunc: twilio.Provider}

	if debugMode {
		err := plugin.Debug(context.Background(), "github.com/mpon/terraform-provider-twilio", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
	plugin.Serve(opts)

}
