package main

import (
	_ "github.com/moby/swarmkit/v2/protobuf/plugin/authenticatedwrapper"
	_ "github.com/moby/swarmkit/v2/protobuf/plugin/deepcopy"
	_ "github.com/moby/swarmkit/v2/protobuf/plugin/raftproxy"
	_ "github.com/moby/swarmkit/v2/protobuf/plugin/storeobject"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	for _, opt := range []func(*descriptor.FileDescriptorProto){
		vanity.TurnOffGoGettersAll,
		vanity.TurnOffGoStringerAll,
		vanity.TurnOnMarshalerAll,
		vanity.TurnOnStringerAll,
		vanity.TurnOnUnmarshalerAll,
		vanity.TurnOnSizerAll,
		vanity.TurnOffGoUnrecognizedAll,
		vanity.TurnOffGoUnkeyedAll,
		vanity.TurnOffGoUnrecognizedAll,
		vanity.TurnOffGoSizecacheAll,
		CustomNameID,
	} {
		vanity.ForEachFile(files, opt)
	}

	resp := command.Generate(req)
	command.Write(resp)
}
