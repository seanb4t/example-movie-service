package internal

import (
	"github.com/rs/zerolog"
)

const GinContextKey = "ginContext"

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
