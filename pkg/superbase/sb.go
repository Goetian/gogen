package supabase

import (
	"errors"
	"os"

	"github.com/nedpals/supabase-go"
)

var (
	Client    *supabase.Client
	sbHost    string
	sbSecrete string
)

func Init() error {

	sbHost = os.Getenv("SUPABASE_URL")
	if sbHost == "" {
		return errors.New("supabase host missing")
	}

	sbSecrete = os.Getenv("SUPABASE_SECRET")

	if sbSecrete == "" {
		return errors.New("supabase secrete missing")
	}

	Client = supabase.CreateClient(sbHost, sbSecrete)

	return nil
}
