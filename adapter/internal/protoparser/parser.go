package protoparser

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

func ParseFile(fnames []string, paths []string) ([]*desc.FileDescriptor, error) {
	paths = append(paths, ".")

	p := &protoparse.Parser{
		ImportPaths: paths,
	}
	return p.ParseFiles(fnames...)
}
