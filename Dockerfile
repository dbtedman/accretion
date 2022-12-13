FROM node:18.12.1-bullseye as build_node

RUN npm install -g pnpm@7.18.1

COPY ./ /code

WORKDIR /code

RUN make install_pnpm build_ts


FROM golang:1.19.3-bullseye as build_go

COPY ./ /code

WORKDIR /code

COPY --from=build_node /code/ui/dist/ /code/ui/dist/

RUN make install_go build_go


FROM scratch

COPY --from=build_go /code/bin/accretion /usr/local/bin/accretion

CMD [ "/usr/local/bin/accretion", "serve" ]
