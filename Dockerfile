FROM alpine:3.12.0

ADD ginGormDemo /
ADD . /

# Add docker-compose-wait tool -------------------
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

WORKDIR /

EXPOSE 88

CMD ["/ginGormDemo"]