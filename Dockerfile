# Production image
FROM alpine:latest AS deploy
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY escapepod .
COPY escapepod.toml.sample escapepod.toml
CMD ["./escapepod", "serve"]
