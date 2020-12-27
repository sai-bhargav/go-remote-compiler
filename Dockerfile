FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update

RUN apt-get install -y ruby

RUN apt-get install -y golang

WORKDIR /app

EXPOSE 5000
