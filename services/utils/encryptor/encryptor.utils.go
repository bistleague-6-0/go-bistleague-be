package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"io"
)

func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(generateKey(passphrase)))
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(generateKey(passphrase)))
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func generateKey(passphrase string) string {
	hasher := sha256.New()
	hasher.Write([]byte(passphrase))
	return hex.EncodeToString(hasher.Sum(nil))[:32] // 32 bytes for AES-256
}

type RefreshKey struct {
	Uid       string           `json:"uid,omitempty"`
	TeamID    string           `json:"team_id,omitempty"`
	ExpiresAt *jwt.NumericDate `json:"exp,omitempty"`
}

func EncryptRefreshKey(myStruct RefreshKey, passphrase string) (string, error) {
	jsonData, err := json.Marshal(myStruct)
	if err != nil {
		return "", err
	}

	encryptedData, err := encrypt(jsonData, passphrase)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(encryptedData), nil
}

func DecryptRefreshKey(encryptedData string, passphrase string) (RefreshKey, error) {
	decodedData, err := hex.DecodeString(encryptedData)
	if err != nil {
		return RefreshKey{}, err
	}

	decryptedData, err := decrypt(decodedData, passphrase)
	if err != nil {
		return RefreshKey{}, err
	}

	var myStruct RefreshKey
	err = json.Unmarshal(decryptedData, &myStruct)
	if err != nil {
		return RefreshKey{}, err
	}

	return myStruct, nil
}
