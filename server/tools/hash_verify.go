package tools

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Hasher struct {
	Salt		string
    Iterations  uint32
    Memory      uint32
    Parallelism uint8
    KeyLength   uint32
}

var (
    ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
    ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func (p *Hasher) GenerateFromPassword(password string) (encodedHash string, err error) {
    hash := argon2.IDKey([]byte(password), []byte(p.Salt), p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

    // Base64 encode the salt and hashed password.
    b64Salt := base64.RawStdEncoding.EncodeToString([]byte(p.Salt))
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)

    // Return a string using the standard encoded hash representation.
    encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

    return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
    // Extract the parameters, salt and derived key from the encoded password
    // hash.
    p, hash, err := decodeHash(encodedHash)
    if err != nil {
        return false, err
    }

    // Derive the key from the other password using the same parameters.
    otherHash := argon2.IDKey([]byte(password), []byte(p.Salt), p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

    // Check that the contents of the hashed passwords are identical. Note
    // that we are using the subtle.ConstantTimeCompare() function for this
    // to help prevent timing attacks.
    if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
        return true, nil
    }
    return false, nil
}

func decodeHash(encodedHash string) (p *Hasher, hash []byte, err error) {
    vals := strings.Split(encodedHash, "$")
    if len(vals) != 6 {
        return nil, nil, ErrInvalidHash
    }

    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, ErrIncompatibleVersion
    }

    p = &Hasher{}
    _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
    if err != nil {
        return nil, nil, err
    }

    salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
    if err != nil {
        return nil, nil, err
    }
    p.Salt = string(salt)

    hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
    if err != nil {
        return nil, nil, err
    }
    p.KeyLength = uint32(len(hash))

    return p, hash, nil
}