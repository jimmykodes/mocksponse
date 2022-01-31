FROM golang:1-buster as builder
WORKDIR /go/src/mocksponse
ADD . .
ENV GOBIN=/
RUN go install .

FROM debian:buster
COPY --from=builder /mocksponse /usr/local/sbin
ADD recipe.yaml /etc/mocksponse/recipe.yaml
CMD ["mocksponse", "-f", "/etc/mocksponse/recipe.yaml"]