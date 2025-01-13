// Package encrypt provides the encrypter for the application.
package encrypt

import (
	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/store"
)

// Type is the type of the encrypt.
type Type string

const (
	// EncryptTypeCBC is the type of the cbc.
	EncryptTypeCBC Type = "CBC"
	// EncryptTypeCFB is the type of the cfb.
	EncryptTypeCFB Type = "CFB"
	// EncryptTypeCTR is the type of the ctr.
	EncryptTypeCTR Type = "CTR"
)

// Encrypter is the interface for the encrypter.
type Encrypter interface {
	Encrypt(plaintext []byte) (string, error)
	Decrypt(ciphertextBase64 string) ([]byte, error)
}

// Container is the container for the encrypter.
type Container map[string]Encrypter

// NewContainerFromConfig creates a new encrypter container.
func NewContainerFromConfig(str store.Store, conf config.ValidEncryptConfig) (Container, error) {
	ec := make(Container)
	var err error
	for _, e := range conf {
		var encrypter Encrypter
		switch e.Type {
		case config.EncryptTypeStaticCBC:
			encrypter, err = NewStaticEncrypter([]byte(e.Key), EncryptTypeCBC)
			if err != nil {
				return nil, err
			}
		case config.EncryptTypeStaticCFB:
			encrypter, err = NewStaticEncrypter([]byte(e.Key), EncryptTypeCFB)
			if err != nil {
				return nil, err
			}
		case config.EncryptTypeStaticCTR:
			encrypter, err = NewStaticEncrypter([]byte(e.Key), EncryptTypeCTR)
			if err != nil {
				return nil, err
			}
		case config.EncryptTypeDynamicCBC:
			encrypter, err = NewDynamicEncrypter(str, e.Store.BucketID, e.Store.Key, EncryptTypeCBC)
			if err != nil {
				return nil, err
			}
		case config.EncryptTypeDynamicCFB:
			encrypter, err = NewDynamicEncrypter(str, e.Store.BucketID, e.Store.Key, EncryptTypeCFB)
			if err != nil {
				return nil, err
			}
		case config.EncryptTypeDynamicCTR:
			encrypter, err = NewDynamicEncrypter(str, e.Store.BucketID, e.Store.Key, EncryptTypeCTR)
			if err != nil {
				return nil, err
			}
		}
		ec[e.ID] = encrypter
	}
	return ec, nil
}
