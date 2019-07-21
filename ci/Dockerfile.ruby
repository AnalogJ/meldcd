ARG PACKAGE_TYPE=ruby

#################################################
#
# Base
# This container should not be used as a runtime environment.
# It is based off a massive build image (crossbuild) which has lots of unnecessary build tools
# It does not actually build the capsulecd executable
# It runs unit tests for each supported engine type.
#
# Use the docker containers in https://github.com/AnalogJ/capsulecd-docker as an example of what a
# proper runtime-environment for CapsuleCD looks like.
#
#################################################

FROM analogj/libgit2-crossbuild:linux-amd64 AS base
MAINTAINER Jason Kulatunga <jason@thesparktree.com>
WORKDIR /go/src/github.com/analogj/capsulecd

RUN apt-get update && apt-get install -y --no-install-recommends \
 	apt-transport-https \
    ca-certificates \
	&& rm -rf /var/lib/apt/lists/* \
	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . .

## download deps & move libgit2 library into expected location.
RUN dep ensure \
	&& mkdir -p /go/src/github.com/analogj/capsulecd/vendor/gopkg.in/libgit2/git2go.v25/vendor/libgit2/build \
	&& cp -r /usr/local/lib/libgit2/lib/pkgconfig/. /go/src/github.com/analogj/capsulecd/vendor/gopkg.in/libgit2/git2go.v25/vendor/libgit2/build/

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

RUN ci/test-build.sh ${PACKAGE_TYPE}

##################################################
##
## Ruby
##
##################################################
FROM ruby:2.4 AS ruby
MAINTAINER Jason Kulatunga <jason@thesparktree.com>

RUN apt-get update && apt-get install -y --no-install-recommends \
 	apt-transport-https \
    ca-certificates \
    git \
    curl \
	&& rm -rf /var/lib/apt/lists/* \
	&& gem install rubocop \
	&& gem install rake \
	&& gem install bundler-audit


# Install GOLANG
ENV GO_VERSION 1.8.3
RUN curl -fsSL "https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz" \
	| tar -xzC /usr/local

ENV PATH="/go/bin:/usr/local/go/bin:${PATH}" \
	GOPATH="/go"

COPY --from=base /go/src/github.com/analogj/capsulecd /go/src/github.com/analogj/capsulecd

WORKDIR /go/src/github.com/analogj/capsulecd

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt

CMD ci/test-coverage.sh
