package godatatools

import (
	"bz.moh.epi/godatatools/auth"
	"bz.moh.epi/godatatools/store"
)

type Server struct {
	DbRepository   store.Store
	BackendBaseURL string
	GoData         auth.GoData
}
