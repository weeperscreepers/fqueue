package fqueue

import (
  "fmt"
  // "log"
);

// the interface we'll be exporting
type FunctionQueue interface {
  Add(func())
  RunNext() bool
  RunAll()
}

/*
  Slice fqueue (abandoned)
*/
// type messageQueue [](func())
//
// func (m *messageQueue) Add (f (func())) {
//   *m = append(*m, f);
// }
//
// func (m *messageQueue) RunNext() {
//   var f (func())
//   f,*m = (*m)[0], (*m)[1:]
//   f()
// }

// A convenient alias
type funcChannel chan (func())

// Asynchronously add a function to the channel
func (m *funcChannel) Add (f (func())) {
  go func(){ (*m) <- f } ()
}

// Channels are FIFO, run the first one added, synchronously.
func (m *funcChannel) RunNext() bool {
  fmt.Println("running next", m)
  select {
    case  f := <- (*m):
        fmt.Println("really running")
        f();
        return true;
    default:
      return false;
  }
}

// Run all functions in the channel one after the other
func (m *funcChannel) RunAll() {
  finished := m.RunNext();
  for finished {
    finished = m.RunNext();
  }
}

// Create a new FunctionQueue using a channel
func NewFunctionQueue() FunctionQueue {
  funcQ := make( funcChannel );
  return &funcQ;
}
