FROM golang:1.23.2-bookworm

ARG USER=user
ARG UID=1000
ARG GID=$UID

RUN groupadd --gid $GID $USER \
    && useradd --uid $UID --gid $GID -m $USER \
    && apt-get update \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir /app \
    && chown $UID:$GID /app

WORKDIR /app

USER $USER

COPY --chown=$UID:$GID go.mod go.sum ./

RUN go mod download \
    && go install github.com/air-verse/air@latest

COPY --chown=$UID:$GID . ./

CMD [ "air" ]
