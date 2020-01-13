package hedera

import (
	"bytes"
	"crypto/ed25519"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPrivateKeyStr = "302e020100300506032b657004220420db484b828e64b2d8f12ce3c0a0e93a0b8cce7af1bb8f39c97732394482538e10"

const testPublicKeyStr = "302a300506032b6570032100e0c8ec2758a5879ffac226a13c0c516b799e72e35141a0dd828f94d37988a4b7"

// generated by hedera-keygen-java, not used anywhere
const testMnemonic = "inmate flip alley wear offer often piece magnet surge toddler submit right radio absent pear floor belt raven price stove replace reduce plate home"
const testMnemonicKey = "302e020100300506032b657004220420853f15aecd22706b105da1d709b4ac05b4906170c2b9c7495dff9af49e1391da"

// backup phrase generated by the iOS wallet, not used anywhere
const iosMnemonicString = "tiny denial casual grass skull spare awkward indoor ethics dash enough flavor good daughter early hard rug staff capable swallow raise flavor empty angle"

// private key for "default account", should be index 0
const iosDefaultPrivateKey = "5f66a51931e8c99089472e0d70516b6272b94dd772b967f8221e1077f966dbda2b60cf7ee8cf10ecd5a076bffad9a7c7b97df370ad758c0f1dd4ef738e04ceb6"

// backup phrase generated by the Android wallet, also not used anywhere
const androidMnemonicString = "ramp april job flavor surround pyramid fish sea good know blame gate village viable include mixed term draft among monitor swear swing novel track"

// private key for "default account", should be index 0
const androidDefaultPrivateKey = "c284c25b3a1458b59423bc289e83703b125c8eefec4d5aa1b393c2beb9f2bae66188a344ba75c43918ab12fa2ea4a92960eca029a2320d8c6a1c3b94e06c9985"

func TestEd25519PrivateKeyGenerate(t *testing.T) {
	key, err := GenerateEd25519PrivateKey()

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(key.String(), ed25519PrivateKeyPrefix))
}

func TestEd25519PrivateKeyExternalSerialization(t *testing.T) {
	key, err := Ed25519PrivateKeyFromString(testPrivateKeyStr)

	assert.NoError(t, err)
	assert.Equal(t, testPrivateKeyStr, key.String())
}

func TestEd25519PrivateKeyExternalSerializationForConcatenatedHex(t *testing.T) {
	keyStr := "db484b828e64b2d8f12ce3c0a0e93a0b8cce7af1bb8f39c97732394482538e10e0c8ec2758a5879ffac226a13c0c516b799e72e35141a0dd828f94d37988a4b7"
	key, err := Ed25519PrivateKeyFromString(keyStr)

	assert.NoError(t, err)
	assert.Equal(t, testPrivateKeyStr, key.String())
}

func TestEd25519PrivateKeyExternalSerializationForRawHex(t *testing.T) {
	keyStr := "db484b828e64b2d8f12ce3c0a0e93a0b8cce7af1bb8f39c97732394482538e10"
	key, err := Ed25519PrivateKeyFromString(keyStr)

	assert.NoError(t, err)
	assert.Equal(t, testPrivateKeyStr, key.String())
}

func TestEd25519PublicKeyExternalSerializationForDerEncodedHex(t *testing.T) {
	key, err := Ed25519PublicKeyFromString(testPublicKeyStr)

	assert.NoError(t, err)
	assert.Equal(t, testPublicKeyStr, key.String())
}

func TestEd25519PublicKeyExternalSerializationForRawHex(t *testing.T) {
	keyStr := "e0c8ec2758a5879ffac226a13c0c516b799e72e35141a0dd828f94d37988a4b7"
	key, err := Ed25519PublicKeyFromString(keyStr)

	assert.NoError(t, err)
	assert.Equal(t, testPublicKeyStr, key.String())
}

func TestEd25519PrivateKeyFromMnemonic(t *testing.T) {
	mnemonic, err := MnemonicFromString(testMnemonic)
	assert.NoError(t, err)

	key, err := Ed25519PrivateKeyFromMnemonic(mnemonic, "")

	assert.NoError(t, err)
	assert.Equal(t, testMnemonicKey, key.String())
}

func TestIOSPrivateKeyFromMnemonic(t *testing.T) {
	mnemonic, err := MnemonicFromString(iosMnemonicString)
	assert.NoError(t, err)

	key, err := Ed25519PrivateKeyFromMnemonic(mnemonic, "")
	assert.NoError(t, err)

	derivedKey, err := key.Derive(0)
	assert.NoError(t, err)

	expectedKey, err := Ed25519PrivateKeyFromString(iosDefaultPrivateKey)
	assert.NoError(t, err)

	assert.Equal(t, expectedKey.keyData, derivedKey.keyData)
}

func TestAndroidPrivateKeyFromMnemonic(t *testing.T) {
	mnemonic, err := MnemonicFromString(androidMnemonicString)
	assert.NoError(t, err)

	key, err := Ed25519PrivateKeyFromMnemonic(mnemonic, "")
	assert.NoError(t, err)

	derivedKey, err := key.Derive(0)
	assert.NoError(t, err)

	expectedKey, err := Ed25519PrivateKeyFromString(androidDefaultPrivateKey)
	assert.NoError(t, err)

	assert.Equal(t, expectedKey.keyData, derivedKey.keyData)
}

func TestSigning(t *testing.T) {
	priKey, err := Ed25519PrivateKeyFromString(testPrivateKeyStr)
	pubKey, err := Ed25519PublicKeyFromString(testPublicKeyStr)
	testSignData := []byte("this is the test data to sign")
	signature := priKey.Sign(testSignData)

	assert.NoError(t, err)
	assert.True(t, ed25519.Verify(pubKey.Bytes(), []byte("this is the test data to sign"), signature))
}

func TestGeneratedMnemonicToWorkingPrivateKey(t *testing.T) {
	mnemonic, err := GenerateMnemonic()

	assert.NoError(t, err)

	privateKey, err := mnemonic.ToPrivateKey("")

	assert.NoError(t, err)

	message := []byte("this is a test message")

	signature := privateKey.Sign(message)

	assert.True(t, ed25519.Verify(privateKey.PublicKey().Bytes(), message, signature))
}

func TestEd25519PrivateKeyFromKeystore(t *testing.T) {
	privatekey, err := Ed25519PrivateKeyFromKeystore([]byte(testKeystore), passphrase)
	assert.NoError(t, err)

	actualPrivateKey, err := Ed25519PrivateKeyFromString(testKeystoreKeyString)
	assert.NoError(t, err)

	assert.Equal(t, actualPrivateKey.keyData, privatekey.keyData)
}

func TestEd25519PrivateKey_Keystore(t *testing.T) {
	privateKey, err := Ed25519PrivateKeyFromString(testPrivateKeyStr)
	assert.NoError(t, err)

	keystore, err := privateKey.Keystore(passphrase)
	assert.NoError(t, err)

	ksPrivateKey, err := parseKeystore(keystore, passphrase)
	assert.NoError(t, err)

	assert.Equal(t, privateKey.keyData, ksPrivateKey.keyData)
}

func TestEd25519PrivateKey_ReadKeystore(t *testing.T) {
	actualPrivateKey, err := Ed25519PrivateKeyFromString(testKeystoreKeyString)
	assert.NoError(t, err)

	keystoreReader := bytes.NewReader([]byte(testKeystore))

	privateKey, err := Ed25519PrivateKeyReadKeystore(keystoreReader, passphrase)
	assert.NoError(t, err)

	assert.Equal(t, actualPrivateKey.keyData, privateKey.keyData)
}
