FROM golang
MAINTAINER gphper
WORKDIR /home/ginadmin/
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' > /etc/timezone
EXPOSE 20010
CMD ["tail","-f","/dev/null"]