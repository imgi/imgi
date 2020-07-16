FROM golang:1.14

ENV DEBIAN_FRONTEND=noninteractive

ARG VIPS_VERSION=8.9.2

# Sources mirror
# COPY sources.list /etc/apt/

# Missing dependencies: PDFium ImageMagick pangoft2 libniftiio
# OpenEXR OpenSlide libheif
RUN apt update && apt install --no-install-recommends -y \
    curl ca-certificates \
    build-essential pkg-config \
    libglib2.0-dev libexpat1-dev libjpeg62-turbo-dev \
    libexif-dev libgif-dev librsvg2-dev libpoppler-glib-dev \
    libgsf-1-dev libtiff5-dev libfftw3-dev liblcms2-dev \
    libpng-dev libimagequant-dev liborc-0.4-dev libmatio-dev \
    libcfitsio-dev libwebp-dev

RUN curl -OL https://github.com/libvips/libvips/releases/download/v${VIPS_VERSION}/vips-${VIPS_VERSION}.tar.gz \
    && tar -zxf vips-${VIPS_VERSION}.tar.gz && cd vips-${VIPS_VERSION} \
    && ./configure --disable-gtk-doc --disable-gtk-doc-html --disable-gtk-doc-pdf \
    && make && make install && ldconfig

WORKDIR ${GOPATH}/src/github.com/imgi/imgi
COPY . .

RUN go build -o imgi cmd/imgi/main.go
