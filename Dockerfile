FROM ubuntu:18.04

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get install -y gcc
RUN apt-get install -y g++
RUN apt-get install -y python3
RUN apt-get install -t python2
RUN apt-get install -y golang-go	
