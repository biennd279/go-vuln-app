package domain

import "github.com/google/uuid"

var (
	EmptyUID = uid{}
)

type UID interface {
	String() string
	id()
}

type uid struct {
	value uuid.UUID
}

type UIDGenerator interface {
	NewUID() UID
	UIDFromString(s string) (UID, error)
}

type UIDGeneratorImpl struct{}

var i UIDGenerator = &UIDGeneratorImpl{}

func NewUID() UID {
	return i.NewUID()
}

func UIDFromString(s string) (UID, error) {
	return i.UIDFromString(s)
}

func (u uid) id() {
	//TODO implement me
	panic("implement me")
}

func (u uid) String() string {
	return u.value.String()
}

// IsEmptyUID is empty uid
func IsEmptyUID(uid UID) bool {
	return uid == EmptyUID
}

func (g *UIDGeneratorImpl) NewUID() UID {
	return uid{
		value: uuid.New(),
	}
}

func (g *UIDGeneratorImpl) UIDFromString(s string) (UID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uid{}, err
	}
	return uid{value: id}, nil
}
