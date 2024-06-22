FROM registry.hub.docker.com/library/golang:1.22.4-bookworm as build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go list -e $(go list -f '{{.Path}}' -m all)

COPY main.go .
COPY pkg pkg

RUN go install


FROM registry.hub.docker.com/library/python:3.12.4-bookworm

WORKDIR /solver

COPY problems problems
COPY evaluateShared.py evaluateShared.py
COPY --from=build /go/bin/vpr solver
RUN ls

CMD ["python3", "evaluateShared.py", "--cmd" , "/solver/solver", "--problemDir", "problems"]


