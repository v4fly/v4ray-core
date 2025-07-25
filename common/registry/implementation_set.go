package registry

import (
	"github.com/golang/protobuf/proto"

	"github.com/v4fly/v4ray-core/v0/common/protoext"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

type implementationSet struct {
	AliasLookup map[string]*implementation
}

type CustomLoader func(data []byte, loader LoadByAlias) (proto.Message, error)

type implementation struct {
	FullName string
	Alias    []string
	Loader   CustomLoader
}

func (i *implementationSet) RegisterImplementation(name string, opt *protoext.MessageOpt, loader CustomLoader) {
	alias := opt.GetShortName()

	impl := &implementation{
		FullName: name,
		Alias:    alias,
	}

	for _, aliasName := range alias {
		i.AliasLookup[aliasName] = impl
	}
}

func (i *implementationSet) findImplementationByAlias(alias string) (string, CustomLoader, error) {
	impl, found := i.AliasLookup[alias]
	if found {
		return impl.FullName, impl.Loader, nil
	}
	return "", nil, newError("cannot find implementation by alias: ", alias)
}

func newImplementationSet() *implementationSet {
	return &implementationSet{AliasLookup: map[string]*implementation{}}
}
