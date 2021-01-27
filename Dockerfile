FROM golang:1.15-buster

RUN apt-get update
RUN apt-get install -y \
   apt-transport-https \
   ca-certificates \
   curl \
   gnupg-agent  \
   software-properties-common \
   gettext-base \
   yamllint python3-pkg-resources shellcheck unzip

RUN curl -fsSL https://get.docker.com | sh

RUN wget -O /usr/bin/yq https://github.com/mikefarah/yq/releases/download/2.1.2/yq_linux_amd64
RUN chmod 555 /usr/bin/yq

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/"$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)"/bin/linux/amd64/kubectl
RUN chmod +x kubectl && mv kubectl /usr/local/bin/

RUN curl -LO https://git.io/get_helm.sh
RUN chmod 700 get_helm.sh
RUN ./get_helm.sh

ARG PROTOC_VER=3.8.0
ARG PROTOC_ZIP=protoc-$PROTOC_VER-linux-x86_64.zip
RUN curl -OL https://github.com/google/protobuf/releases/download/v$PROTOC_VER/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN rm -f $PROTOC_ZIP

COPY . /go/src/github.com/cisco-app-networking/networkservicemesh
WORKDIR /go/src/github.com/cisco-app-networking/networkservicemesh

#RUN yamllint -c .yamllint.yml $(git ls-files '*.yaml' '*.yml')

RUN go mod download

RUN go install k8s.io/code-generator/cmd/deepcopy-gen
RUN go install github.com/golang/protobuf/protoc-gen-go
RUN go get golang.org/x/tools/cmd/stringer
RUN go get github.com/networkservicemesh/cloudtest@v0.2
RUN GO111MODULE="on" go get -u sigs.k8s.io/kind@v0.8.1

RUN go generate ./...
#RUN make vet check
RUN go build ./...

#RUN CGO_ENABLED=0 GOOS=linux go build -o ./cloudtest ./test/cloudtest/cmd/cloudtest
CMD cloudtest

#RUN go test ./test/integration --list .* -tags basic recover usecase

