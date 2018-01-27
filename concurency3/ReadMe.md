select -
select can have a timeout statement.  case <-time.After(25 * time.Second):
select will block until one is selected.
default will pass through.  It will exit from select. 
select can receive channel value.  case value1 := <- channel1 : 
select can send channel value.  case channel2 <- value2 :

Error handler
Error, Panic, Recover
type error interface {
    Error() string
}