package yotsuba

import (
    "bytes"
    "compress/gzip"
    "crypto/md5"
    "encoding/base64"
    "fmt"
    "hash"
    "io"
    "log"
)

// Enigma - Wrapper to Cryptographer and Hasher
//
// Basic go-to collection for cryptography and encoding.
type Enigma struct {
    h hash.Hash
}

func (self *Enigma) Hasher() hash.Hash {
    if self.h == nil {
        self.h = md5.New()
    }

    self.h.Reset()

    return self.h
}

func (self *Enigma) Hash(Content []byte) string {
    sum := self.Hasher().Sum(Content)

    return fmt.Sprintf("%x", sum)
}

func (self *Enigma) HashString(Content string) string {
    data := []byte(Content)

    return self.Hash(data)
}

func (self *Enigma) Compress(content []byte) []byte {
    b := new(bytes.Buffer)
    w := gzip.NewWriter(b)

    defer w.Close()

    w.Write(content)

    return b.Bytes()
}

func (self *Enigma) Decompress(content []byte) []byte {
    readingB := new(bytes.Buffer)
    writingB := new(bytes.Buffer)

    readingB.Write(content)

    r, _ := gzip.NewReader(readingB)

    defer r.Close()

    io.Copy(writingB, r)

    return writingB.Bytes()
}

func (self *Enigma) B64decode(content string) string {
    rawMessage, err := base64.StdEncoding.DecodeString(content)

    if err != nil {
        log.Fatal("tori.enigma.Enigma.B64decode/error:", err)
    }

    return string(rawMessage)
}
