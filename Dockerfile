# pipa/go dockerfile
# alpine Go...
# @date 03/2018
FROM golang

MAINTAINER Luis Matute

ARG app_env
ENV APP_ENV $app_env
ENV PORT 8888
ENV WEBROOT /go/src/github.com/pipa/app

WORKDIR ${WEBROOT}
ADD ./app ${WEBROOT}

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
  then \
  app; \
  else \
  go get github.com/pilu/fresh && \
  fresh; \
  fi

EXPOSE ${PORT}
