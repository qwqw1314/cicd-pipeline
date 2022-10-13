FROM alpine
COPY cicd-pipeline /home/
CMD ["/home/cicd-pipeline"]