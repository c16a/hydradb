package internal

import "github.com/spf13/afero"

type Config struct {
	Server  *Server  `json:"server"`
	Storage *Storage `json:"storage"`
}

type Server struct {
	Address string `json:"address"`
}

type Storage struct {
	Directory  string   `json:"directory"`
	Filesystem afero.Fs `json:"-"`
}
