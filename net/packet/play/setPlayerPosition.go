package play

import "github.com/dynamitemc/aether/net/io"

// serverbound
const PacketIdSetPlayerPosition = 0x1A

type SetPlayerPosition struct {
	X, Y, Z  float64
	OnGround bool
}

func (SetPlayerPosition) ID() int32 {
	return 0x1A
}

func (s *SetPlayerPosition) Encode(w io.Writer) error {
	if err := w.Double(s.X); err != nil {
		return err
	}
	if err := w.Double(s.Y); err != nil {
		return err
	}
	if err := w.Double(s.Z); err != nil {
		return err
	}
	return w.Bool(s.OnGround)
}

func (s *SetPlayerPosition) Decode(r io.Reader) error {
	if err := r.Double(&s.X); err != nil {
		return err
	}
	if err := r.Double(&s.Y); err != nil {
		return err
	}
	if err := r.Double(&s.Z); err != nil {
		return err
	}
	return r.Bool(&s.OnGround)
}