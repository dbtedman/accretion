FROM node:18.12.1-bullseye@sha256:e9ad817b0d42b4d177a4bef8a0aff97c352468a008c3fdb2b4a82533425480df as build_node

RUN npm install -g pnpm@7.18.1

COPY ./ /code

WORKDIR /code

RUN make install_pnpm build_ts


FROM golang:1.19.3-bullseye@sha256:74fe5785cff148dfb88053d1caa8011baccdc8ad0616af77baa124438f218caa as build_go

COPY ./ /code

WORKDIR /code

COPY --from=build_node /code/ui/dist/ /code/ui/dist/

RUN make install_go build_go


FROM scratch

COPY --from=build_go /code/bin/accretion /usr/local/bin/accretion

CMD [ "/usr/local/bin/accretion", "serve" ]
