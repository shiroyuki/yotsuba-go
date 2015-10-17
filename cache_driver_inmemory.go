package yotsuba

// In-memory Cache Driver
type InMemoryCacheDriver struct {
    CacheDriver // implements yotsuba.CacheDriver

    Enigma      *Enigma
    Compressed  bool
    MemoryTable map[string][]byte
}

func NewInMemoryCacheDriver(enigma *Enigma, compressed bool) *InMemoryCacheDriver {
    imcd := InMemoryCacheDriver{
        Enigma:      enigma,
        Compressed:  compressed,
        MemoryTable: make(map[string][]byte),
    }

    return &imcd
}

func (self *InMemoryCacheDriver) Load(key string) []byte {
    content, existed := self.MemoryTable[key]

    if !existed {
        return nil
    }

    if !self.Compressed {
        return content
    }

    return (*self.Enigma).Decompress(content)
}

func (self *InMemoryCacheDriver) Save(key string, content []byte) {
    if self.MemoryTable == nil {
        self.MemoryTable = make(map[string][]byte)
    }

    if !self.Compressed {
        self.MemoryTable[key] = content

        return
    }

    compressed           := (*self.Enigma).Compress(content)
    self.MemoryTable[key] = compressed
}
