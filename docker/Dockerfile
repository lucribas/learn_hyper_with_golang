FROM golang:1.19

RUN apt-get update
RUN apt-get -y install git curl vim iproute2 procps lsb-release

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /opt/app

RUN go install -v github.com/cweill/gotests/gotests@latest && \
go install -v github.com/fatih/gomodifytags@latest && \
go install -v github.com/josharian/impl@latest && \
go install -v github.com/haya14busa/goplay/cmd/goplay@latest && \
go install -v github.com/go-delve/delve/cmd/dlv@latest && \
go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
go install -v golang.org/x/tools/gopls@latest

ARG USERNAME=user
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
        && useradd --uid $USER_UID --gid $USER_GID -ms /bin/bash $USERNAME \
        #
        # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
        && apt-get update \
        && apt-get install -y sudo \
        && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
        && chmod 0440 /etc/sudoers.d/$USERNAME

RUN groupmod --gid $USER_GID $USERNAME \
    && usermod --uid $USER_UID --gid $USER_GID $USERNAME \
    && chown -R $USER_UID:$USER_GID /home/$USERNAME
RUN chown -R $USERNAME /opt/app

USER $USERNAME
WORKDIR /opt/app

CMD ["/bin/bash"]
