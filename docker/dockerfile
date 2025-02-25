FROM centos:8
COPY install.sh /
RUN sh install.sh
ENV PATH "$PATH:/usr/local/go/bin"
# ENV GOPATH "/root/go"
# ENV MONGODB_URI "mongodb://root:password@172.23.0.3:27017/?retryWrites=true&w=majority"
# EXPOSE 8080
# EXPOSE 80
# EXPOSE 22