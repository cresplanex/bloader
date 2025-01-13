package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/cresplanex/bloader/internal/utils"
)

// Encrypt encrypts plaintext using the specified encryption method with the given key and IV.
func Encrypt(plaintext, key []byte, encryptMethod Type) (string, error) {
	var ciphertext []byte
	var iv []byte

	// Choose encryption method
	switch encryptMethod {
	case EncryptTypeCBC:
		block, err := aes.NewCipher(key)
		if err != nil {
			return "", fmt.Errorf("failed to create cipher block: %w", err)
		}
		iv, err = utils.GenerateRandomBytes(aes.BlockSize)
		if err != nil {
			return "", fmt.Errorf("failed to generate IV: %w", err)
		}
		plaintext = PKCS7Padding(plaintext, aes.BlockSize)
		ciphertext = make([]byte, len(plaintext))
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ciphertext, plaintext)
	case EncryptTypeCFB:
		block, err := aes.NewCipher(key)
		if err != nil {
			return "", fmt.Errorf("failed to create cipher block: %w", err)
		}
		iv, err = utils.GenerateRandomBytes(aes.BlockSize)
		if err != nil {
			return "", fmt.Errorf("failed to generate IV: %w", err)
		}
		ciphertext = make([]byte, len(plaintext))
		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(ciphertext, plaintext)
	case EncryptTypeCTR:
		block, err := aes.NewCipher(key)
		if err != nil {
			return "", fmt.Errorf("failed to create cipher block: %w", err)
		}
		iv, err = utils.GenerateRandomBytes(aes.BlockSize)
		if err != nil {
			return "", fmt.Errorf("failed to generate IV: %w", err)
		}
		ciphertext = make([]byte, len(plaintext))
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(ciphertext, plaintext)
	default:
		return "", fmt.Errorf("unsupported encryption method: %v", encryptMethod)
	}

	// Prepend IV to ciphertext
	combined := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

// Decrypt decrypts ciphertext using the specified encryption method with the given key.
func Decrypt(ciphertextBase64 string, key []byte, encryptMethod Type) ([]byte, error) {
	combined, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 ciphertext: %w", err)
	}

	var plaintext []byte
	var iv []byte
	var ciphertext []byte

	// Extract IV and ciphertext
	switch encryptMethod {
	case EncryptTypeCBC, EncryptTypeCFB, EncryptTypeCTR:
		if len(combined) < aes.BlockSize {
			return nil, fmt.Errorf("invalid ciphertext length")
		}
		iv = combined[:aes.BlockSize]
		ciphertext = combined[aes.BlockSize:]
	default:
		return nil, fmt.Errorf("unsupported decryption method: %v", encryptMethod)
	}

	// Choose decryption method
	switch encryptMethod {
	case EncryptTypeCBC:
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, fmt.Errorf("failed to create cipher block: %w", err)
		}
		plaintext = make([]byte, len(ciphertext))
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(plaintext, ciphertext)
		plaintext = PKCS7Unpadding(plaintext)
	case EncryptTypeCFB:
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, fmt.Errorf("failed to create cipher block: %w", err)
		}
		plaintext = make([]byte, len(ciphertext))
		stream := cipher.NewCFBDecrypter(block, iv)
		stream.XORKeyStream(plaintext, ciphertext)
	case EncryptTypeCTR:
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, fmt.Errorf("failed to create cipher block: %w", err)
		}
		plaintext = make([]byte, len(ciphertext))
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(plaintext, ciphertext)
	default:
		return nil, fmt.Errorf("unsupported decryption method: %s", encryptMethod)
	}

	return plaintext, nil
}

// PKCS7Unpadding removes padding from the plaintext.
func PKCS7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// PKCS7Padding adds padding to the plaintext to make it a multiple of the block size.
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
