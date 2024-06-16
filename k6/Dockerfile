FROM grafana/k6:latest
COPY ./script.js ./script.js
COPY ./run.sh ./run.sh
ENTRYPOINT [ "/bin/sh", "run.sh" ]
