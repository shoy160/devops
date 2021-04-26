FROM alpine
WORKDIR /publish
ADD fed/dist/. dist
ADD template/. template
ADD devops .
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT [ "./devops" ]