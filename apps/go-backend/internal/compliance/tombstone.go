package compliance

type Tombstone struct {
	EntityID      string
	ReplacementID string
	Reason        string
}

func ApplyTombstone(entityID string, reason string) Tombstone {
	return Tombstone{
		EntityID:      entityID,
		ReplacementID: "tombstone:" + entityID,
		Reason:        reason,
	}
}
