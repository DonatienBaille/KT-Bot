# We have to get the last certificates to be able to connect to Discord
# or SSL verification will fail
FROM alpine:3.19.0 as alpine
RUN apk add --no-cache \
  ca-certificates

FROM scratch
COPY --from=alpine \
  /etc/ssl/certs/ca-certificates.crt \
  /etc/ssl/certs/ca-certificates.crt
COPY kt-bot /
ENTRYPOINT [ "/kt-bot" ]