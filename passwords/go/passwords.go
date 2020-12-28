package passwords

import (
	"strings"
)

type PasswordCandidate struct {
	Min      int
	Max      int
	Letter   string
	Passowrd string
}

func (p *PasswordCandidate) IsValid() bool {
	occurencies := strings.Count(p.Passowrd, p.Letter)
	return occurencies >= p.Min && occurencies <= p.Max
}
