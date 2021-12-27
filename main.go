package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/picatz/protoc-gen-go-validate/pkg/validate/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gp *protogen.Plugin) error {
		gp.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		gen := gen.NewGenerator(gp)

		return gen.Generate()
	})
}
