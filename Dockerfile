FROM debian:buster-slim

LABEL maintainer="dortiz@devops.faith"

RUN apt-get update && \
	apt-get install -y ca-certificates && \
	update-ca-certificates && \
	rm -rf /var/lib/apt/lists/*

ADD krakend /usr/bin/krakend

RUN useradd -r -c "KrakenD user" -U krakend

USER krakend

VOLUME [ "/etc/krakend" ]

WORKDIR /etc/krakend

COPY ./krakend.json /etc/krakend/krakend.json
# COPY ./plugins/ /etc/krakend/plugin/
# COPY ./plugins/ralali-auth-interceptor/*.so /etc/krakend/plugin/
COPY ./plugins/ralali-header-transformator/*.so /etc/krakend/plugin/


ENTRYPOINT [ "/usr/bin/krakend" ]
CMD [ "run", "-c", "/etc/krakend/krakend.json" ]

EXPOSE 8000 8090
