package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

// Wallet contains the user's public and private keys
type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// GenerateKeys generates a new RSA private and public key pair
func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// CreateWallet creates a new wallet with a new set of RSA keys
func CreateWallet() (*Wallet, error) {
	privateKey, publicKey, err := GenerateKeys()
	if err != nil {
		return nil, err
	}

	return &Wallet{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

// SignTransaction signs a transaction using the wallet's private key
func (wallet *Wallet) SignTransaction(tx *Transaction) (string, error) {
	data := fmt.Sprintf("%s%s%f%t", tx.Sender, tx.Receiver, tx.Amount, tx.Coinbase)
	hash := sha256.Sum256([]byte(data))

	signature, err := rsa.SignPKCS1v15(rand.Reader, wallet.PrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifyTransaction verifies a transaction's signature using the public key
func VerifyTransaction(tx *Transaction, publicKey *rsa.PublicKey, signature string) error {
	data := fmt.Sprintf("%s%s%f%t", tx.Sender, tx.Receiver, tx.Amount, tx.Coinbase)
	hash := sha256.Sum256([]byte(data))

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signatureBytes)
	if err != nil {
		return errors.New("Transaction Signature not valid.")
	}
	return nil
}
