FROM golang:1.19.3-bullseye as build

COPY ./ /code

WORKDIR /code

RUN make install_go build_go


FROM scratch

COPY --from=build /code/bin/accretion /usr/local/bin/accretion

CMD [ "/usr/local/bin/accretion", "serve" ]
