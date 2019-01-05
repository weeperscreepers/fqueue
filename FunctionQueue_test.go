package fqueue

import (
  "testing"
  "fmt"
  "time"
  //"github.com/satori/go.uuid"
)

func line() {
  fmt.Println("---------------------------")
}

func TestFunctionQueue(t *testing.T){
    line()
    fmt.Println("Beginning Function Queue Test")
    line()
    hello := func(){ fmt.Println("hello") };
    // hola := func(){ fmt.Println("hola") };
    goodbye := func(){ fmt.Println("goodbye") };
    adios := func(){ fmt.Println("adios") }

    // mq := messageQueue{}
    // fmt.Println("goign to add...")
    // mq.Add( hello )
    // mq.Add( hola )
    // mq.RunNext()
    // mq.RunNext()

    skiplength := 100*time.Millisecond;

    fmt.Println("channel to add...")
    cmq := make( funcChannel )

    cmq.Add( goodbye );
    cmq.Add( adios );

    time.Sleep(skiplength)

    cmq.RunNext()
    cmq.RunNext()

    time.Sleep(skiplength)

    fq := NewFunctionQueue();
    fq.Add( hello );
    fq.Add( adios );

    time.Sleep(skiplength)

    fq.RunNext()
    fq.RunNext()

    time.Sleep(skiplength)
}


func TestQueueExhaustion(t *testing.T){
  line()
  fmt.Println("Beginning Function Queue Exhaustion Test")
  line()
  hello := func(){ fmt.Println("hello") };
  hola := func(){ fmt.Println("hola") };
  goodbye := func(){ fmt.Println("goodbye") };
  // adios := func(){ fmt.Println("adios") }

  skiplength := 100*time.Millisecond;

  fq := NewFunctionQueue();
  fq.Add(hello)
  fq.Add(hola)
  time.Sleep(skiplength)
  fq.RunNext();

  fq.Add(goodbye)
  time.Sleep(skiplength)
  fq.RunAll()

  fq.Add(hello)
  fq.Add(hola)
  time.Sleep(skiplength)
  fq.RunAll()

  // for { fmt.Println("runnnnen ... ")}


}
