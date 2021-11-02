FROM ubuntu

WORKDIR /

COPY --chown=root:root ./.docker/app/supermarket /supermarket

EXPOSE 3000

USER root:root

ENTRYPOINT [ "/supermarket" ]