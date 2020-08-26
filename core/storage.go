package core

import (
	"github.com/gabek/owncast/config"
	"github.com/gabek/owncast/core/ffmpeg"
	"github.com/gabek/owncast/core/storageproviders"
)

var (
	usingExternalStorage = false
)

func setupStorage() error {
	handler = ffmpeg.HLSHandler{}
	handler.Storage = _storage
	fileWriter.SetupFileWriterReceiverService(&handler)

	if config.Config.S3.Enabled {
		_storage = &storageproviders.S3Storage{}
	} else {
		_storage = &storageproviders.LocalStorage{}
	}

	if err := _storage.Setup(); err != nil {
		return err
	}

	return nil
}
