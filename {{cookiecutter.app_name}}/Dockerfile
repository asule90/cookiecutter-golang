# Build Stage
FROM {{cookiecutter.docker_base_image}}:{{cookiecutter.docker_base_image_version}} AS build-stage
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Final Stage
FROM gcr.io/distroless/base-debian11 as production
COPY --from=build-stage /app/main /
EXPOSE 8003

ENTRYPOINT [ "./main" ]