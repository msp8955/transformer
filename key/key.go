package key

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

func Read(reader *bytes.Buffer, toRead ...interface{}) error {
	for _, data := range toRead {
		switch value := data.(type) {
		case *[]byte:
			decoded, err := reader.ReadBytes(0)
			if err != nil {
				return err
			}
			*value = decoded[:len(decoded)-1]
		case *string:
			var decoded []byte
			if err := Read(reader, &decoded); err != nil {
				return err
			}
			*value = string(decoded)
		case *uint8:
			var decoded uint8
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			*value = decoded
		case *int8:
			var decoded uint8
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			if decoded >= math.MaxUint8-math.MaxInt8 {
				*value = int8(decoded - (math.MaxUint8 - math.MaxInt8))
			} else {
				*value = int8(decoded) - int8(math.MaxUint8-math.MaxInt8-1) - 1
			}
		case *uint16:
			var decoded uint16
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			*value = decoded
		case *int16:
			var decoded uint16
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			if decoded >= math.MaxUint16-math.MaxInt16 {
				*value = int16(decoded - (math.MaxUint16 - math.MaxInt16))
			} else {
				*value = int16(decoded) - int16(math.MaxUint16-math.MaxInt16-1) - 1
			}
		case *uint32:
			var decoded uint32
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			*value = decoded
		case *int32:
			var decoded uint32
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			if decoded >= math.MaxUint32-math.MaxInt32 {
				*value = int32(decoded - (math.MaxUint32 - math.MaxInt32))
			} else {
				*value = int32(decoded) - int32(math.MaxUint32-math.MaxInt32-1) - 1
			}
		case *uint64:
			var decoded uint64
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			*value = decoded
		case *int64:
			var decoded uint64
			if err := binary.Read(reader, binary.BigEndian, &decoded); err != nil {
				return err
			}
			if decoded >= math.MaxUint64-math.MaxInt64 {
				*value = int64(decoded - (math.MaxUint64 - math.MaxInt64))
			} else {
				*value = int64(decoded) - int64(math.MaxUint64-math.MaxInt64-1) - 1
			}
		}
	}
	return nil
}

func Write(writer io.Writer, toWrite ...interface{}) error {
	for _, data := range toWrite {
		switch value := data.(type) {
		case []byte:
			if bytes.Contains(value, []byte{'\x00'}) {
				return fmt.Errorf("Cannot encode embedded null characters")
			}
			if _, err := writer.Write(append(value, '\x00')); err != nil {
				return err
			}
		case string:
			if err := Write(writer, []byte(value)); err != nil {
				return err
			}
		case uint8:
			if err := binary.Write(writer, binary.BigEndian, value); err != nil {
				return err
			}
		case int8:
			var uint8value uint8
			if value >= 0 {
				uint8value = uint8(value)+(math.MaxUint8-math.MaxInt8)
			} else {
				uint8value = uint8(value-math.MinInt8)
			}
			if err := Write(writer, uint8value); err != nil {
				return err
			}
		case uint16:
			if err := binary.Write(writer, binary.BigEndian, value); err != nil {
				return err
			}
		case int16:
			var uint16value uint16
			if value >= 0 {
				uint16value = uint16(value)+(math.MaxUint16-math.MaxInt16)
			} else {
				uint16value = uint16(value-math.MinInt16)
			}
			if err := Write(writer, uint16value); err != nil {
				return err
			}
		case uint32:
			if err := binary.Write(writer, binary.BigEndian, value); err != nil {
				return err
			}
		case int32:
			var uint32value uint32
			if value >= 0 {
				uint32value = uint32(value)+(math.MaxUint32-math.MaxInt32)
			} else {
				uint32value = uint32(value-math.MinInt32)
			}
			if err := Write(writer, uint32value); err != nil {
				return err
			}
		case uint64:
			if err := binary.Write(writer, binary.BigEndian, value); err != nil {
				return err
			}
		case int64:
			var uint64value uint64
			if value >= 0 {
				uint64value = uint64(value)+(math.MaxUint64-math.MaxInt64)
			} else {
				uint64value = uint64(value-math.MinInt64)
			}
			if err := Write(writer, uint64value); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Lexicographic encoding not available for type %T", value)
		}
	}
	return nil
}