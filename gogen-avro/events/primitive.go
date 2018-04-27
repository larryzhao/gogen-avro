// Code generated by github.com/alanctgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     foo.avsc
 */

package events

import (
	"fmt"
	"io"
	"math"
)

type ByteWriter interface {
	Grow(int)
	WriteByte(byte) error
}

type StringWriter interface {
	WriteString(string) (int, error)
}

func encodeInt(w io.Writer, byteCount int, encoded uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	// To avoid reallocations, grow capacity to the largest possible size
	// for this integer
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}

	if encoded == 0 {
		if bw != nil {
			err = bw.WriteByte(0)
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(0))
		}
	} else {
		for encoded > 0 {
			b := byte(encoded & 127)
			encoded = encoded >> 7
			if !(encoded == 0) {
				b |= 128
			}
			if bw != nil {
				err = bw.WriteByte(b)
				if err != nil {
					return err
				}
			} else {
				bb = append(bb, b)
			}
		}
	}
	if bw == nil {
		_, err := w.Write(bb)
		return err
	}
	return nil

}

func readFoo(r io.Reader) (*Foo, error) {
	var str = &Foo{}
	var err error
	str.Bar, err = readUnionStringNull(r)
	if err != nil {
		return nil, err
	}
	str.UserId, err = readInt(r)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func readInt(r io.Reader) (int32, error) {
	var v int
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		v |= int(b&127) << shift
		if b&128 == 0 {
			break
		}
	}
	datum := (int32(v>>1) ^ -int32(v&1))
	return datum, nil
}

func readLong(r io.Reader) (int64, error) {
	var v uint64
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		v |= uint64(b&127) << shift
		if b&128 == 0 {
			break
		}
	}
	datum := (int64(v>>1) ^ -int64(v&1))
	return datum, nil
}

func readNull(_ io.Reader) (interface{}, error) {
	return nil, nil
}

func readString(r io.Reader) (string, error) {
	len, err := readLong(r)
	if err != nil {
		return "", err
	}

	// makeslice can fail depending on available memory.
	// We arbitrarily limit string size to sane default (~2.2GB).
	if len < 0 || len > math.MaxInt32 {
		return "", fmt.Errorf("string length out of range: %d", len)
	}

	bb := make([]byte, len)
	_, err = io.ReadFull(r, bb)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}

func readUnionStringNull(r io.Reader) (UnionStringNull, error) {
	field, err := readLong(r)
	var unionStr UnionStringNull
	if err != nil {
		return unionStr, err
	}
	unionStr.UnionType = UnionStringNullTypeEnum(field)
	switch unionStr.UnionType {
	case UnionStringNullTypeEnumString:
		val, err := readString(r)
		if err != nil {
			return unionStr, err
		}
		unionStr.String = val
	case UnionStringNullTypeEnumNull:
		val, err := readNull(r)
		if err != nil {
			return unionStr, err
		}
		unionStr.Null = val

	default:
		return unionStr, fmt.Errorf("Invalid value for UnionStringNull")
	}
	return unionStr, nil
}

func writeFoo(r *Foo, w io.Writer) error {
	var err error
	err = writeUnionStringNull(r.Bar, w)
	if err != nil {
		return err
	}
	err = writeInt(r.UserId, w)
	if err != nil {
		return err
	}

	return nil
}

func writeInt(r int32, w io.Writer) error {
	downShift := uint32(31)
	encoded := uint64((uint32(r) << 1) ^ uint32(r>>downShift))
	const maxByteSize = 5
	return encodeInt(w, maxByteSize, encoded)
}

func writeLong(r int64, w io.Writer) error {
	downShift := uint64(63)
	encoded := uint64((r << 1) ^ (r >> downShift))
	const maxByteSize = 10
	return encodeInt(w, maxByteSize, encoded)
}

func writeNull(_ interface{}, _ io.Writer) error {
	return nil
}

func writeString(r string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil {
		return err
	}
	if sw, ok := w.(StringWriter); ok {
		_, err = sw.WriteString(r)
	} else {
		_, err = w.Write([]byte(r))
	}
	return err
}

func writeUnionStringNull(r UnionStringNull, w io.Writer) error {
	err := writeLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionStringNullTypeEnumString:
		return writeString(r.String, w)
	case UnionStringNullTypeEnumNull:
		return writeNull(r.Null, w)

	}
	return fmt.Errorf("Invalid value for UnionStringNull")
}
