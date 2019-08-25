ARG NODE_VERSION=12.8.0
ARG GOLANG_VERSION=1.12.9

# Frontend Dependencies (cache)
FROM node:${NODE_VERSION} as frontend-deps
WORKDIR /escapepod
ADD frontend-vue/yarn.lock frontend-vue/package.json frontend-vue/
RUN cd frontend-vue && yarn install

# Frontend build
FROM node:${NODE_VERSION} as frontend
WORKDIR /escapepod
COPY . .
COPY --from=frontend-deps /escapepod/frontend-vue/node_modules frontend-vue/node_modules
RUN make build-frontend

# Backend build
FROM golang:${GOLANG_VERSION} as build
RUN go get github.com/knadh/stuffbin/stuffbin
WORKDIR /escapepod
COPY . .
COPY --from=frontend /escapepod/frontend-vue/dist frontend-vue/dist
RUN make dist

# Production image
FROM alpine:latest AS deploy
RUN apk --no-cache add ca-certificates
WORKDIR /escapepod
COPY --from=build /escapepod/escapepod .
COPY escapepod.toml.sample escapepod.toml
CMD ["./escapepod", "serve"]
