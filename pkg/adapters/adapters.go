package adapters

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/adapters/filesystemadapter"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/adapters/inmemoryadapter"
)

var (
	NewInMemoryAdapter   = inmemoryadapter.NewInMemoryAdapter
	NewFileSystemAdapter = filesystemadapter.NewFileSystemAdapter
)
