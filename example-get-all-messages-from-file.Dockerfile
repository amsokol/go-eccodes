FROM amsokol/go-eccodes/centos-builder:latest

RUN go get github.com/amsokol/go-eccodes/cmd/example-get-all-messages-from-file

ADD ./test-data/ARPEGE_0.1_SP1_00H12H_201709290000.grib2 data.grib2

CMD ["example-get-all-messages-from-file", "-file", "./data.grib2"]
