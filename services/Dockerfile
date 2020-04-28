FROM busybox as base
WORKDIR /app
EXPOSE 5050

FROM golang:1.13.4 as build

WORKDIR /src
COPY . .
RUN chmod +x ./build.sh
RUN ./build.sh
RUN mkdir -p /app/zoneinfo
RUN mv ./myblogs /app
RUN mv ./app_data /app
RUN ls ./zoneinfo
RUN mv ./zoneinfo/zoneinfo.zip /app/zoneinfo/zoneinfo.zip

FROM base as final
WORKDIR /app
COPY --from=build /app .
RUN ls .
VOLUME [ "/app/app_data" ]
ENV ZONEINFO=/app/zoneinfo/zoneinfo.zip
ENV APP_MODE=release
# ENV GIN_MODE=release
ENTRYPOINT [ "./myblogs" ]