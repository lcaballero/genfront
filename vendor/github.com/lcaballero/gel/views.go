package gel

// Views represent a list of views that can be rendered as a one View.
type Views interface {
	Viewable
	View
	Add(views ...View) Views
}

// Fragment is the concrete type of the Views interface.
type Fragment struct {
	views []View
}

// NewViews allocates an empty views list to which move views can be added.
func NewFragment() *Fragment {
	return &Fragment{
		views: make([]View, 0),
	}
}

// Add appends the given views to the currently established list of views.
func (v *Fragment) Add(views ...View) Views {
	v.views = append(v.views, views...)
	return v
}

// ToViews implements to the Viewable interface returning a concrete Node
// fragment instance.
func (v *Fragment) ToView() View {
	return Frag(v.views...)
}

// ToNode implements the View interface, producing the final transformation
// to a *Node which can be rendered to stream.
func (v Fragment) ToNode() *Node {
	return v.ToView().ToNode()
}

// Len returns the number of views held by the Fragment.
func (v Fragment) Len() int {
	return len(v.views)
}

// Viewable is the interface for anything that can be rendered into a View.
type Viewable interface {
	ToView() View
}

// ToView is the functional equivalent of Viewable, which allows an
// empty function returning a View to be a Viewable.
type ToView func() View

// ToView method implments the Viewable interface for the ToView func.
func (v ToView) ToView() View {
	return v()
}

// View interface represents anything that can be turned into a Node via the
// ToNode() function.
type View interface {
	ToNode() *Node
}

// ToNode is an empty function that returns a rendered/completed Node.
type ToNode func() *Node

// ToNode here implements the View interface, meaning that a ToNode function
// can now be used like a View.
func (t ToNode) ToNode() *Node {
	return t()
}
