FROM golang:1.14

RUN go get github.com/cespare/reflex

RUN apt-get update -y && apt-get install -y jq

RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') && \
    curl -o /usr/local/bin/swagger -L'#' "$download_url" && \
    chmod +x /usr/local/bin/swagger

# make the go module cache folder writable by all users otherwise linux users will not be able to write to it
RUN mkdir -p /go/pkg/mod/cache && chmod a+rw -R /go/pkg/mod/

COPY reflex.conf /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
