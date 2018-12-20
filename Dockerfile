FROM ubuntu:18.04

RUN apt-get update && apt-get install -y --no-install-recommends apt-utils
RUN apt-get -y upgrade
RUN apt-get install -y gcc
RUN apt-get install -y g++
RUN apt-get install -y python3
RUN apt-get install -y golang-go	
