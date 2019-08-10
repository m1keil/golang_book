package ex9_1

/*
Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program. The result should indicate whether the
transaction succeeded or failed due to insufficient funds. The message sent to the monitor goroutine must contain both
the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back to Withdraw.
*/

// withdrawal request
type request struct {
	amount int
	response chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan request)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraw <- request{amount, ch}

	return <- ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount

		case balances <- balance:

		case request := <- withdraw:
			if balance - request.amount >= 0 {
				balance -= request.amount
				request.response <- true
			} else {
				request.response <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
