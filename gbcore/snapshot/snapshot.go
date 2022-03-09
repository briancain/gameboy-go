package gbcore

import "time"

type Snapshot struct {
	id int

	Time time.Time

	Parent *Snapshot
}

func (s *Snapshot) TakeSnapshot(parent *Snapshot) (*Snapshot, error) {
	return nil, nil
}
