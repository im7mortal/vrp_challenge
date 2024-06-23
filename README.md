# Instructions 

To run follow run sections. Algorith explained in the Algorith section.

![Screenshot from 2024-06-23 14-35-39](https://github.com/im7mortal/vrp_challenge/assets/5336231/34afa20f-b16c-4305-bf3b-6c8ce62b91e1)


## Run

The fastest way to run it. Is use [Dockerfile](Dockerfile). 
```shell
 docker build --build-arg dir="Training Problems" --build-arg pyScript="evaluateShared.py" -t vpr:latest .
 docker run vpr:latest
```

The directory must contain the problems directory and evaluateShared.py. You can set them in build arguments `dir` and `pyScript` accordingly.
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

This allowed achieving a mean cost of 52041.86997702847 and an average runtime of 4.130744934082031ms. In the next iteration, I decided to tune it with some versions N recursive brute force Salesman algorithm.

This resulted in a slight improvement in the mean cost, down to 49508.511593292424, but a significant increase in runtime to 8068.529975414276ms. I figured out that some test cases contained a multitude of small vectors close to the depot (especially problems 5 and 6). What caused very deep recursions and consequently, longer runtimes.

So I decided to apply following KNOWHOW : If the algorithm's execution exceeded 1 second for N=3 (branching), the best result from N=1,2 would be selected instead. 

Currently, further improvements are being pursued through experimentation with the Random Evaluation function.


```
// 1 52041.86997702847 <-- NEAREST NEIGHBOR
// 2 49628.99589848467
// 3 49324.77775315616
// 4 49384.56281527733
// 5 49408.26447408905
```

## Code

I am trying to demonstrate proficiency in GoLang and in general programming.
