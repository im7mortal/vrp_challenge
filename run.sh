 docker build --build-arg dir="Training Problems" --build-arg pyScript="evaluateShared.py" -t vpr:latest . && \
 docker run vpr:latest