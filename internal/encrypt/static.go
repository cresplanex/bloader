package encrypt

// StaticEncrypter is the static encrypter.
type StaticEncrypter struct {
	key    []byte
	method Type
}

// NewStaticEncrypter creates a new static encrypter.
func NewStaticEncrypter(key []byte, method Type) (*StaticEncrypter, error) {
	return &StaticEncrypter{
		key:    key,
		method: method,
	}, nil
}

// Encrypt encrypts the plaintext using the static encrypter.
func (e *StaticEncrypter) Encrypt(plaintext []byte) (string, error) {
	ciphertext, err := Encrypt(plaintext, e.key, e.method)
	if err != nil {
		return "", err
	}
	return ciphertext, nil
}

// Decrypt decrypts the ciphertext using the static encrypter.
func (e *StaticEncrypter) Decrypt(ciphertextBase64 string) ([]byte, error) {
	plaintext, err := Decrypt(ciphertextBase64, e.key, e.method)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

var _ Encrypter = (*StaticEncrypter)(nil)
