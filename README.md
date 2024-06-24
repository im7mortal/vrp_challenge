# Instructions 

To run follow run sections. Algorithm explained in the [Algorithm](Algorithm) section.

![Screenshot from 2024-06-23 14-35-39](https://github.com/im7mortal/vrp_challenge/assets/5336231/34afa20f-b16c-4305-bf3b-6c8ce62b91e1)


## Run

The fastest way to run it. Is use [Dockerfile](Dockerfile). 
```shell
 docker build --build-arg dir="Training Problems" --build-arg pyScript="evaluateShared.py" -t vpr:latest .
 docker run vpr:latest
```

Or run [run.sh](run.sh)

The directory **must contain** `the problems directory` and `evaluateShared.py`. You can set them in build arguments `dir` and `pyScript` accordingly.
Default values 

```dockerfile
ARG dir="Training Problems"
ARG pyScript="evaluateShared.py"
```


Alternatively, you can compile the Go binary for your operating system using the command:
```shell
go build .
```

## Algorithm

The original intent was to implement the A* or smt similar. However, I started with the Nearest Neighbor algorithm.

This allowed achieving a mean cost of **52041.86997702847** and an average runtime of **4.130744934082031ms**. In the next iteration, I decided to tune it with some versions N recursive brute force Salesman algorithm.

This resulted in a slight improvement in the mean cost, down to 49508.511593292424, but a significant increase in runtime to 8068.529975414276ms. I figured out that some test cases contained a multitude of small vectors close to the depot (especially problems 5 and 6). What caused very deep recursions and consequently, longer runtimes.

So I decided to apply following KNOWHOW : If the algorithm's execution exceeded 1 second for N=3 (branching), the best result from N=1,2 would be selected instead. 

Later I applied different evaluators functions.

### Evaluators

After we have a list of possible routes we start evaluate them. 

1. GetTheBestByLengthAndCostMin - we filter the longest routes and choose route with minimal cost
2. GetTheBestByLengthAndCostMax - we filter the longest routes and choose route with maximum cost
3. GetTheBestByLengthAndRandom - we filter the longest routes and then choose random one from it. To ensure determinacy,
    [I seed the random generator with content of source file. For every iteration we recreate a random generator so it always give the same random sequence.](https://github.com/im7mortal/vrp_challenge/blob/main/pkg/solvers/utils/parser.go#L81-L108). I use CHACHA8 fast random generator which was added in rand/v2 go22.


See [performance](PERFORMANCE_LOG.md) for details.
```
// 1 52041.86997702847 <-- NEAREST NEIGHBOR
// 2 49628.99589848467
// 3 49324.77775315616
...

// 48341.15786902501  <-- run N = 1, 2, 3 with Min, Max, Random evaluators 
```

## Code

I am trying to demonstrate proficiency in GoLang and in general programming.

Most golish part is this [clumsy code](https://github.com/im7mortal/vrp_challenge/blob/0f8ef3a515d5668c157e8b1731f31d6a844bd8bc/pkg/solvers/parrallel.go#L45-L119). As I mentioned in the Algorithm part. Some test cases can proceed for less than 1ms with N=1 and 150 seconds for N=3.
    We start all tasks in the same time in different gorutines. I decided do not implement process limiter as it's 9 gorutines maximum for current implementation and 3 of them finish almost immediately.
    What matters , it's to intercept the test cases significant amount of time and cancel them immediately. Following logic handle it.
