FROM kong:3.4

USER root

RUN apt-get update && apt-get install -y gettext-base && rm -rf /var/lib/apt/lists/*

COPY plugins/token-validator /usr/local/share/lua/5.1/kong/plugins/token-validator
COPY docker-entrypoint.sh /docker-entrypoint-custom.sh
RUN chmod +x /docker-entrypoint-custom.sh

USER kong

ENTRYPOINT ["/docker-entrypoint-custom.sh"]

