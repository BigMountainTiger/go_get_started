FROM ubuntu:22.10

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

EXPOSE 80
WORKDIR /app

COPY .out/* .

ENTRYPOINT [ "./start"]
CMD []