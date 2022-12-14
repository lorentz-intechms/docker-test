FROM ubuntu:20.04

WORKDIR ./code
RUN apt-get update

COPY . /code

CMD echo ${APOLLO_URL} >> test.txt