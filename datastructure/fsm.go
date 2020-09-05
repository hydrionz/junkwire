package datastructure

import "fmt"

//State 状态
type State struct {
	name  string
	trans map[string]trans //key是event value是state的name
}

//Event trans event
type Event struct {
	Name       string
	ConnPacket interface{}
}

type trans struct {
	name   string      //to state
	action func(Event) //action after trans
}

//Fsm 状态机
type Fsm struct {
	states  map[string]*State //key是state的name value是state指针
	current string            //currentEvent
}

//NewFsm 创建
func NewFsm(init string) *Fsm {
	m := new(Fsm)
	m.current = init
	m.states = make(map[string]*State)
	return m
}

//AddRule 添加规则
func (m *Fsm) AddRule(state string, event Event, toState string, action func(Event)) {
	if m.states[state] == nil {
		m.states[state] = new(State)
		m.states[state].name = state
		m.states[state].trans = make(map[string]trans)
	}
	m.states[state].trans[event.Name] = trans{name: toState, action: action}
}

//OnEvent 事件发生的回调
func (m *Fsm) OnEvent(event Event) {
	currentState := m.states[m.current]
	nextTrans, ok := currentState.trans[event.Name]
	if !ok {
		eventName := event.Name
		if eventName == "" {
			eventName = "nil"
		}
		fmt.Printf("stata:%s has no event %s\n", currentState.name, eventName)
		return
	}
	nextStateName := nextTrans.name
	m.current = nextStateName
	nextTrans.action(event)
}
