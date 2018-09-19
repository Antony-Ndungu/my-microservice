FROM scratch
LABEL maintainer="Antony Ndungu <antony.ndungu.maina@gmail.com"

EXPOSE 8080

COPY ./server ./

ENTRYPOINT ["./server"]
