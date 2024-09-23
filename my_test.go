package iso8583

import (
	"encoding/hex"
	"testing"

	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/prefix"
)

func FuzzCustomBCD(f *testing.F) {
	spec := &field.Spec{
		Description: "my field",
		Enc:         encoding.BCD,
		Pref:        prefix.Binary.LL,
		Length:      99999,
	}

	f.Add([]byte("0"))
	f.Add([]byte("01"))
	f.Add([]byte("012"))
	f.Add([]byte("0123"))
	f.Add([]byte("01234"))
	f.Add([]byte("012345"))
	f.Add([]byte("0123456"))

	f.Fuzz(func(t *testing.T, rawData []byte) {
		f := field.NewBinary(spec)

		// rawData := []byte("123")

		err := f.SetBytes(rawData)
		if err != nil {
			t.Errorf("failed to set byte: %v", err)
		}

		packed, err := f.Pack()
		if err != nil {
			t.Errorf("failed to pack the field: %v", err)
		}

		t.Logf("Packed message: %v", hex.EncodeToString(packed))

		f2 := field.NewBinary(spec)
		_, err = f2.Unpack(packed)
		if err != nil {
			t.Errorf("failed to unpack: %v", err)
		}

		// if len(rawData) != read {
		// 	t.Errorf("the read length does not equal the raw BCD character count: %v != %v", read, len(rawData))
		// }

		unpacked, err := f2.Bytes()
		if err != nil {
			t.Errorf("failed to read bytes from the unpacked field: %v", err)
		}

		t.Logf("Unpacked data: %v", string(unpacked))
	})
}

func FuzzCustomASCIIHex(f *testing.F) {
	spec := &field.Spec{
		Description: "my field",
		Enc:         encoding.ASCIIHexToBytes,
		Pref:        prefix.Binary.LL,
		Length:      99999,
	}

	// f.Add([]byte("0"))
	f.Add([]byte("01"))
	// f.Add([]byte("012"))
	// f.Add([]byte("0123"))
	// f.Add([]byte("01234"))
	// f.Add([]byte("012345"))
	// f.Add([]byte("0123456"))

	f.Fuzz(func(t *testing.T, rawData []byte) {
		f := field.NewBinary(spec)

		// rawData := []byte("123")

		err := f.SetBytes(rawData)
		if err != nil {
			t.Errorf("failed to set byte: %v", err)
		}

		packed, err := f.Pack()
		if err != nil {
			t.Errorf("failed to pack the field: %v", err)
		}

		t.Logf("Packed message: %v", hex.EncodeToString(packed))

		f2 := field.NewBinary(spec)
		_, err = f2.Unpack(packed)
		if err != nil {
			t.Errorf("failed to unpack: %v", err)
		}

		// if len(rawData) != read {
		// 	t.Errorf("the read length does not equal the raw BCD character count: %v != %v", read, len(rawData))
		// }

		unpacked, err := f2.Bytes()
		if err != nil {
			t.Errorf("failed to read bytes from the unpacked field: %v", err)
		}

		t.Logf("Unpacked data: %v", string(unpacked))
	})
}
