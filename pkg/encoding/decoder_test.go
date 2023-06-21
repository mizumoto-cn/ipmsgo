package encoding_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	zh_cn "golang.org/x/text/encoding/simplifiedchinese"

	. "github.com/mizumoto-cn/ipmsgo/pkg/encoding"
)

type GB18030Decoder struct{}

func (d *GB18030Decoder) Decode(src []byte, ctx context.Context) ([]byte, error) {
	return zh_cn.GB18030.NewDecoder().Bytes(src)
}

const (
	GB18030    = "gb18030"
	HELLOWORLD = "hello, 世界 。！"
)

func TestDecoderAll(t *testing.T) {
	decoded := HELLOWORLD
	// UTF-8: 0x68 0x65 0x6c 0x6c 0x6f 0x2c 0x20 0xe4 0xb8 0x96 0xe7 0x95 0x8c 0x20 0xe3 0x80 0x82 0xef 0xbc 0x81

	utf8Bytes := []byte{
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0x20, 0xe3, 0x80, 0x82, 0xef, 0xbc, 0x81,
	}

	// GBK: 68 65 6C 6C 6F 2C 20 CA C0 BD E7 20 A1 A3 A3 A1

	GBKBytes := []byte{
		0x68, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0xCA, 0xC0, 0xBD, 0xE7, 0x20, 0xA1, 0xA3, 0xA3, 0xA1,
	}

	// GB18030: 68 65 6C 6C 6F 2C 20 CA C0 BD E7 20 A1 A3 A3 A1

	GB18030Bytes := []byte{
		0x68, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0xCA, 0xC0, 0xBD, 0xE7, 0x20, 0xA1, 0xA3, 0xA3, 0xA1,
	}

	// ShiftJIS: 68 65 6c 6c 6f 2c 20 90 A2 8A 45 20 81 42 81 49

	ShiftJISBytes := []byte{
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x90, 0xA2, 0x8A, 0x45, 0x20, 0x81, 0x42, 0x81, 0x49,
	}

	dMap := DecoderFactoryMap{}

	cases := []struct {
		name    string
		f       DecoderFactoryFunc
		encoded []byte
		decoded string
	}{
		{
			name: UTF8,
			f: DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
				return &UTF8Decoder{}, nil
			}),
			encoded: utf8Bytes,
			decoded: decoded,
		},
		{
			name: GBK,
			f: DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
				return &GBKDecoder{}, nil
			}),
			encoded: GBKBytes,
			decoded: decoded,
		},
		{
			name: GB18030,
			f: DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
				return &GB18030Decoder{}, nil
			}),
			encoded: GB18030Bytes,
			decoded: decoded,
		},
		{
			name: ShiftJIS,
			f: DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
				return &ShiftJISDecoder{}, nil
			}),
			encoded: ShiftJISBytes,
			decoded: decoded,
		},
	}
	for _, c := range cases {
		dMap.Register(
			c.name,
			c.f,
		)
		decoder, err := dMap.CreateDecoder(c.name, context.Background())
		assert.Nil(t, err)
		cDecoded, err := decoder.Decode(c.encoded, context.Background())
		assert.Nil(t, err)
		assert.Equal(t, c.decoded, string(cDecoded))
	}
}

func TestCreateFail(t *testing.T) {
	dMap := DecoderFactoryMap{}
	decoder, err := dMap.CreateDecoder("invalid", context.Background())
	assert.Nil(t, decoder)
	assert.NotNil(t, err)
}
