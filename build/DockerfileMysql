FROM mysql:5.7

#设置免密登录
ENV MYSQL_ALLOW_EMPTY_PASSWORD yes

#将所需文件放到容器中
COPY ./build/setup.sh /home/setup.sh
COPY ./build/init.sql /home/init.sql

RUN sed -i 's/\r$//' /home/setup.sh

#设置容器启动时执行的命令
CMD ["sh", "/home/setup.sh"]