FROM bin.private.zooplus.net/platform/jenkins/agent-default

MAINTAINER MadGyvers Team (Customer Fidelization and Segmentation) <team_dev_madgyvers@zooplus.com>

ENV GOLANG_VERSION=1.11.1

ENV ZOOPLUS_ENV DEV

EXPOSE 5222

RUN curl -fsSL https://bin.private.zooplus.net/artifactory/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz -o golang.tar.gz \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /root
ENV GO_HOME /usr/lib/golang
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY lmwtfy/ /usr/local/lmwtfy

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" "$GOPATH/pkg"

#download go modules
CMD cd /usr/local/lmwtfy && go mod tidy && go build cmd/main.go