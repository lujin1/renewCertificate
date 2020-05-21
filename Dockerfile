FROM harbor.arfa.wise-paas.com/lujin/pipeline-node:v1.0.0.1
MAINTAINER Lu jin

COPY renewCertificate /
COPY kubectl-cert_manager /usr/local/bin/kubectl-cert_manager
RUN apk add jq && rm /var/cache/apk/*
CMD ["/renewCertificate"]