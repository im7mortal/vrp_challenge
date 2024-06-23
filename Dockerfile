FROM registry.hub.docker.com/library/golang:1.22.4-bookworm as build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN #go list -e $(go list -f '{{.Path}}' -m all); exit 0

COPY main.go .
COPY pkg pkg

RUN go install


FROM registry.hub.docker.com/library/python:3.12.4-bookworm

WORKDIR /solver

ARG dir="Training Problems"
ARG pyScript="evaluateShared.py"

COPY ${dir} problems
COPY ${pyScript} run.py
COPY --from=build /go/bin/vpr solver

CMD ["python3", "run.py", "--cmd" , "/solver/solver", "--problemDir", "problems"]


