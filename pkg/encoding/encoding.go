package encoding

import (
	"context"
)

var (
	decoderFactoryMap = DecoderFactoryMap{}
)

const (
	UTF8     = "utf8"
	GBK      = "gbk"
	ShiftJIS = "shiftjis"
)

func init() {
	decoderFactoryMap.Register(UTF8, DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
		return &UTF8Decoder{}, nil
	}))
	decoderFactoryMap.Register(GBK, DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
		return &GBKDecoder{}, nil
	}))
	decoderFactoryMap.Register(ShiftJIS, DecoderFactoryFunc(func(ctx context.Context) (Decoder, error) {
		return &ShiftJISDecoder{}, nil
	}))
}

// Decode decodes the given byte slice using the specified decoder.
func Decode(name string, src []byte, ctx context.Context) ([]byte, error) {
	decoder, err := CreateDecoder(name, ctx)
	if err != nil {
		return nil, err
	}
	return decoder.Decode(src, ctx)
}

// Register adds a new decoder factory to the DecoderFactoryMap.
func RegisterDecoderFactory(name string, factory DecoderFactory) {
	decoderFactoryMap.Register(name, factory)
}

// CreateDecoder creates a new decoder based on the given name and context.
func CreateDecoder(name string, ctx context.Context) (Decoder, error) {
	return decoderFactoryMap.CreateDecoder(name, ctx)
}
