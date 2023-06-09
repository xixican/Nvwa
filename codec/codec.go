package codec

type Codec interface {
	Decode(b []byte)
	Encode() []byte
}
