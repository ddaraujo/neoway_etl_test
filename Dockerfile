#Define GOLANG as default base image
FROM golang:latest

#Set Workdir
RUN mkdir /app
ADD . /app
WORKDIR /app

# Fetch and build app
#RUN go get github.com/EBKopec/etl_base_teste
#RUN cd /build && git clone https://github.com/EBKopec/etl_base_teste.git
#RUN cd /build/etl_base_teste/ && go build

# Expose app port 
EXPOSE 8088

# Run service entrypoint
#ENTRYPOINT ["/app/neoway_etl_test"]




