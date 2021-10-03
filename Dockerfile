FROM alpine:latest
WORKDIR service
COPY eprescription-reg-service /service/
EXPOSE 8081
CMD ["/service/eprescription-reg-service"]