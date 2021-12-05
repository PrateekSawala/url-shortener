FROM scratch

ENV URL_RECORD_FILE_PATH /templates/urls.json
ENV TEST_PORT :80

COPY server /server
COPY templates /templates
COPY static /static

CMD ["/server"]