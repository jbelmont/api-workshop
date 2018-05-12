FROM alpine:latest

ENV PORT 8080
EXPOSE 8080

# Add Timezone data package so we can look it up in Go
RUN apk update && apk add tzdata

ADD ./bin/apid /apid

CMD [ "/apid" ]
