FROM harbor.arfa.wise-paas.com/lujin/pipeline-node:v1.0.0.1
MAINTAINER Lu jin

COPY renewCertificate /
RUN apk add jq && rm /var/cache/apk/*

CMD ["/renewCertificate"]