## Task 1 - Matching Behaviour

Take a look at the program matching.go. Explain what happens and why it happens if you make the following changes. Try first to reason about it, and then test your hypothesis by changing and running the program.

* What happens if you remove the `go-command` from the `Seek` call in the `main` function?
    * The program will run only on the main goroutine. Everything will therefore be happen in sequence and in the same order every time, meaning that Anna always sends a message to Bob, Cody to Dave and Eve to nobody. 

* What happens if you switch the declaration `wg := new(sync.WaitGroup)` to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?
    * Using new returns a pointer to the newly created waitgroup which can be passed to other goroutines to share the same waitgroup. Replacing new would instead return the waitgroup itself, which when passed as an argument will generate a copy. The original waitgroup will not care if the 5 copies ever wg.Done().
    
* What happens if you remove the buffer on the channel match?
    * Since the channel would be unbuffered, the sender and receiver in the Seek function would both block until the other is ready, and none of them will be able to proceed, leading to a deadlock.

* What happens if you remove the default-case from the case-statement in the `main` function?
    * Not much will happen in this case as the base case will never be reached since there is an odd number of people. If the number of people instead was even, the case statement would keep trying to read from the channel even after the goroutines finish which would cause the program to never exit. 

## Task 3 - MapReduce

|Variant       | Runtime (ms) | #Samples   |
| ------------ | ------------:| -----------|
| singleworker |      8.31    |     100    |
| mapreduce    |      4.55    |     100    |