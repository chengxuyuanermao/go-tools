package fsm

// Event is the info that get passed as a reference in the callbacks.
type Event struct {
	// FSM is a reference to the current FSM.
	FSM *FSM

	// Event is the event name.
	Event int

	// Src is the state before the transition.
	Src int

	// Dst is the state after the transition.
	Dst int

	// Err is an optional error that can be returned from a callback.
	Err error

	// Args is a optinal list of arguments passed to the callback.
	Args []interface{}
}

// Current returns the current state of the FSM.
func (f *FSM) Current() int {
	//f.stateMu.RLock()
	//defer f.stateMu.RUnlock()
	return f.current
}

type Events []EventDesc

// cKey is a struct key used for keeping the callbacks mapped to a target.
type cKey struct {
	target int

	// callbackType is the situation when the callback will be run.
	callbackType int
}

// eKey is a struct key used for storing the transition map.
type eKey struct {
	// event is the name of the event that the keys refers to.
	event int

	// src is the source from where the event can transition.
	src int
}

type Callback func(*Event)

type FSM struct {
	current int

	// transitions maps events and source states to destination states.
	transitions map[eKey]int

	// callbacks maps events and targers to callback functions.
	callbacks map[cKey]Callback
}

type EventDesc struct {
	// Name is the event name used when calling for a transition.
	Name int

	// Src is a slice of source states that the FSM must be in to perform a
	// state transition.
	Src []int

	// Dst is the destination state that the FSM will be in if the transition
	// succeds.
	Dst int

	Before Callback
	Enter  Callback
	After  Callback
}

const (
	callbackNone int = iota
	callbackBeforeEvent
	//callbackLeaveState
	callbackEnterState
	callbackAfterEvent
)

func NewFSM(initial int, events []EventDesc) *FSM {
	f := &FSM{
		current:     initial,
		transitions: make(map[eKey]int),
		callbacks:   make(map[cKey]Callback),
	}

	for _, e := range events {
		for _, src := range e.Src {
			f.transitions[eKey{e.Name, src}] = e.Dst
		}
		f.callbacks[cKey{e.Name, callbackBeforeEvent}] = e.Before
		f.callbacks[cKey{e.Name, callbackEnterState}] = e.Enter
		f.callbacks[cKey{e.Name, callbackAfterEvent}] = e.After
	}

	return f
}

func NewFSMLudo(initial int, events []*EventDesc) *FSM {
	f := &FSM{
		current:     initial,
		transitions: make(map[eKey]int),
		callbacks:   make(map[cKey]Callback),
	}

	for _, e := range events {
		for _, src := range e.Src {
			f.transitions[eKey{e.Name, src}] = e.Dst
		}
		if e.Before != nil {
			f.callbacks[cKey{e.Name, callbackBeforeEvent}] = e.Before
		}
		if e.Enter != nil {
			f.callbacks[cKey{e.Name, callbackEnterState}] = e.Enter
		}
		if e.After != nil {
			f.callbacks[cKey{e.Name, callbackAfterEvent}] = e.After
		}
	}

	return f
}

func (f *FSM) Event(event int, args ...interface{}) bool {
	if f == nil {
		return false
	}
	dst, ok := f.transitions[eKey{event, f.current}]
	if !ok {
		//log.Error("未找到状态")
	}
	//if dst == event {
	//	log.Error("已切换到该状态")
	//}
	e := &Event{f, event, f.current, dst, nil, args}

	f.afterEventCallbacks(e)

	f.current = event
	f.beforeEventCallbacks(e)
	f.enterStateCallbacks(e)
	return true
}

func (f *FSM) beforeEventCallbacks(e *Event) {
	if fn, ok := f.callbacks[cKey{e.Event, callbackBeforeEvent}]; ok {
		fn(e)
	}
}

func (f *FSM) enterStateCallbacks(e *Event) {
	if fn, ok := f.callbacks[cKey{e.Event, callbackEnterState}]; ok {
		fn(e)
	}
}

func (f *FSM) afterEventCallbacks(e *Event) {
	if fn, ok := f.callbacks[cKey{e.Src, callbackAfterEvent}]; ok {
		fn(e)
	}
}
