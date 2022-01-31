FROM golang:1-buster as builder
WORKDIR /go/src/mock-sponse
ADD . .
ENV GOBIN=/
RUN go install .

FROM debian:buster
COPY --from=builder /mock-sponse /usr/local/sbin
ADD recipe.yaml /etc/mock-sponse/recipe.yaml
CMD ["mock-sponse", "-f", "/etc/mock-sponse/recipe.yaml"]