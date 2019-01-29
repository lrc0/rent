FROM alpine
MAINTAINER ruicai.li@changhong.com
RUN echo "Asia/Shanghai" > /etc/timezones
WORKDIR /workspace
ENV GOPATH=/workspace
RUN apk add --no-cache tzdata
COPY ./conf ./src/rentmanagement/conf
COPY ./swagger-ui ./src/swagger-ui
COPY ./rent ./src/rentmanagement
RUN chmod +x ./src/rentmanagement/rent
ENTRYPOINT ["./src/rentmanagement/rent"]