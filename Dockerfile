#Define GOLANG as default base image
FROM golang:latest

#Set Workdir
RUN mkdir /app
WORKDIR /app

# Fetch and build app
#RUN go get github.com/ddaraujo/neoway_etl_test
RUN git clone https://github.com/ddaraujo/neoway_etl_test.git
RUN cd neoway_etl_test && go build

# Expose app port 
EXPOSE 8888

# Run service entrypoint
ENTRYPOINT ["/app/neoway_etl_test"]
#ENTRYPOINT ["tail"]
#CMD ["-f","/dev/null"]




