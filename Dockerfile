FROM debian:wheezy

LABEL Vincent Bastioni <vbastioni@gmail.com>

# Production apt-get
RUN apt-get -y update\
        && apt-get -y upgrade\
        && apt-get -y install\
            gcc\
            make\
            golang\
            procps

# Development apt-get
RUN apt-get -y install\
            vim\
            git\
            man

COPY ./conf/vimrc /tmp/.vimrc
COPY ./conf/bashrc /tmp/.bashrc
COPY ./conf/bashrc /tmp/prep_env.sh

WORKDIR /home/frozen_server

EXPOSE 3306

COPY src/server src/server
COPY src/lib src/github.com/vbastioni/lib

ENV GOPATH /home/frozen_server
ENV GOBIN /usr/lib/go/bin
ENV PATH $PATH:/usr/lib/go/bin

RUN /bin/rm -f ~/.bashrc\
    && /bin/ln -s /tmp/.vimrc /tmp/.bashrc ~\
    && /bin/rm /tmp/*\
    && /bin/mkdir -p /usr/lib/go/bin

WORKDIR /home/frozen_server/src/server

RUN go build -o /home/frozen_server/server

ENTRYPOINT /home/frozen_server/server --port 3306
