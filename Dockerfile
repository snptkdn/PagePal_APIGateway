FROM golang:1.18.4-alpine3.16 as base
RUN apk update && apk add git

FROM base as develop

# for Vscode extension
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest

# hotreload
RUN go install github.com/cosmtrek/air@latest

FROM base as production

# serverless framework
RUN apk add make
RUN apk add --no-cache nodejs npm
COPY package.json .
RUN npm install -g serverless
