package models

import (
	"github.com/lunixbochs/struc"
	"io"
)

type StrucStream struct {
	Stream  io.ReadWriter
	Options *struc.Options
}

func (s *StrucStream) Pack(vals ...interface{}) error {
	for _, val := range vals {
		err := struc.PackWithOptions(s.Stream, val, s.Options)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StrucStream) Unpack(vals ...interface{}) error {
	for _, val := range vals {
		err := struc.UnpackWithOptions(s.Stream, val, s.Options)
		if err != nil {
			return err
		}
	}
	return nil
}
