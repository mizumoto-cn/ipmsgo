package encoding_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/mizumoto-cn/ipmsgo/pkg/encoding"
)

func TestEncodingAll(t *testing.T) {
	RegisterDecoderFactory(GB18030, DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
		return &GB18030Decoder{}, nil
	}))
	d, err := CreateDecoder(GB18030, context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, d)
	raw := []byte{
		0x68, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0xCA, 0xC0, 0xBD, 0xE7, 0x20, 0xA1, 0xA3, 0xA3, 0xA1,
	}
	decoded, err := d.Decode(raw, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, HELLOWORLD, string(decoded))
	decoded, err = Decode(GB18030, raw, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, HELLOWORLD, string(decoded))
}

func TestInitDecode(t *testing.T) {
	tests := []struct {
		name     string
		encoding string
		input    []byte
		expected []byte
	}{
		{
			name:     "UTF-8",
			encoding: UTF8,
			input:    []byte("hello, world!"),
			expected: []byte("hello, world!"),
		},
		{
			name:     "GBK",
			encoding: GBK,
			input:    []byte{0x68, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0xCA, 0xC0, 0xBD, 0xE7, 0x20, 0xA1, 0xA3, 0xA3, 0xA1},
			expected: []byte(HELLOWORLD),
		},
		{
			name:     "ShiftJIS",
			encoding: ShiftJIS,
			input:    []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x90, 0xA2, 0x8A, 0x45, 0x20, 0x81, 0x42, 0x81, 0x49},
			expected: []byte(HELLOWORLD),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Decode(tt.encoding, tt.input, context.Background())
			if err != nil {
				t.Fatalf("Decode(%q, %v) returned error: %v", tt.encoding, tt.input, err)
			}
			if !bytes.Equal(actual, tt.expected) {
				t.Errorf("Decode(%q, %v) = %v, expected %v", tt.encoding, tt.input, actual, tt.expected)
			}
		})
	}
}

func TestDecodeError(t *testing.T) {
	tests := []struct {
		name    string
		decoder string
		src     []byte
		want    []byte
		wantErr bool
	}{
		{
			name:    "UnknownDecoder",
			decoder: "unknown",
			src:     []byte("hello world"),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.decoder, tt.src, context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
