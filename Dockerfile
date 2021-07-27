############################
# STEP 1 build executable binary
############################
FROM registry.redhat.io/rhel8/go-toolset:1.15 AS builder
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Use go mod
ENV GO111MODULE=on
# Fetch dependencies.
# Using go get requires root.
USER root
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 go build -o /go/bin/edge-api

# Build the migration binary.
RUN CGO_ENABLED=0 go build -o /go/bin/edge-api-migrate cmd/migrate/migrate.go

############################
# STEP 2 build a small image
############################
#FROM registry.redhat.io/ubi8-minimal:latest
FROM quay.io/loadtheaccumulator/ubi8-isotools:latest

COPY --from=builder /go/bin/edge-api /usr/bin
COPY --from=builder /go/bin/edge-api-migrate /usr/bin
COPY --from=builder /src/mypackage/myapp/cmd/spec/openapi.json /var/tmp

# kickstart inject requirements
COPY --from=builder /src/mypackage/myapp/pkg/images/fleetkick.sh /usr/local/bin
COPY --from=builder /src/mypackage/myapp/pkg/images/templateKickstart.ks /usr/local/etc

#RUN microdnf install -y pykickstart mtools xorriso genisoimage syslinux isomd5sum file
ENV MTOOLS_SKIP_CHECK=1

USER 1001

CMD ["edge-api"]
