package filesystemadapter

import (
	"sync"

	artifactmanager "github.com/VictorMarcolino/artifact-manipulator/pkg/adapters"
	"github.com/VictorMarcolino/artifact-manipulator/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
)

type FilesystemAdapter struct {
	ArtifactRepo ports.ArtifactRepositoryI
	mu           sync.RWMutex
}

func NewFileSystemAdapter() *FilesystemAdapter {
	storage := utils.GetPathRelativeToFolder("tmp/data")
	adapter, err := artifactmanager.NewFileSystemArtifactRepository(storage)
	if err != nil {
		panic(err)
	}
	return &FilesystemAdapter{ArtifactRepo: adapter}
}
