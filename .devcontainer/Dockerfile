FROM golang:1.19.3

ARG PACKAGE_LIST="\
    curl \
    gawk \
    git \
    less \
    locales \
    make \
    man-db \
    procps \
    sudo \
    vim \
"

RUN apt-get update \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends ${PACKAGE_LIST}

ARG USERNAME=santa
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create user
RUN groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd -s /bin/bash --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME}

# Give user sudo permissions
RUN echo ${USERNAME} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME}

# Locale generation
RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen && locale-gen

# Tools used by Go extension
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
