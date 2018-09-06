FROM alpine:latest
ADD bin/devops-test bin/devops-test
ENTRYPOINT ["bin/devops-test"]