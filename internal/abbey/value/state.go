package value

type (
	State struct {
		value state
	}

	state interface {
		VisitState(StateVisitor)
	}

	StateVisitor struct {
		Null    func()
		Unknown func()
		Valid   func()
	}

	null    struct{}
	unknown struct{}
	valid   struct{}
)

var (
	_ state = (*null)(nil)
	_ state = (*unknown)(nil)
	_ state = (*valid)(nil)
)

func (null) VisitState(visitor StateVisitor)    { visitor.Null() }
func (unknown) VisitState(visitor StateVisitor) { visitor.Unknown() }
func (valid) VisitState(visitor StateVisitor)   { visitor.Valid() }
func (s State) Visit(visitor StateVisitor)      { s.value.VisitState(visitor) }

func NewNullState() State    { return State{value: null{}} }
func NewUnknownState() State { return State{value: unknown{}} }
func NewValidState() State   { return State{value: valid{}} }
