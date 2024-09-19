package utils

import (
    "bytes"
    "encoding/binary"
    "fmt"
)

func UintToBytes(num interface{}) ([]byte, error) {
    buf := new(bytes.Buffer)
    var err error

    switch v := num.(type) {
    case uint16:
        err = binary.Write(buf, binary.BigEndian, v)
    case uint32:
        err = binary.Write(buf, binary.BigEndian, v)
    case uint64:
        err = binary.Write(buf, binary.BigEndian, v)
    default:
        return nil, fmt.Errorf("unsupported type %T", v)
    }

    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
