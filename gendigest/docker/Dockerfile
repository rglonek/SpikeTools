FROM ubuntu:22.04

LABEL description='Dockerfile for gendigest'

ADD gendigest-amd64 /bin/
ADD gendigest-arm64 /bin/
ADD gendigest /bin/

ENTRYPOINT [ "/bin/gendigest" ]
