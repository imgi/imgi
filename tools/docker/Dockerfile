FROM imgi/imgi-dev AS dev

FROM debian:buster-slim

ENV DEBIAN_FRONTEND=noninteractive

# Sources mirror
# COPY sources.list /etc/apt/

# Missing dependencies: PDFium ImageMagick libniftiio
RUN apt update && apt install --no-install-recommends -y \
    libglib2.0 \
    libexpat1 \
    libjpeg62-turbo \
    libexif12 \
    libgif7 \
    librsvg2-2 \
    libpoppler-glib8 \
    libgsf-1-114 \
    libtiff5 \
    libfftw3-3 \
    liblcms2-2 \
    libpng16-16 \
    libimagequant0 \
    libpango1.0-0 \
    liborc-0.4-0 \
    libmatio4 \
    libcfitsio7 \
    libwebp6 \
    libwebpdemux2 \
    libwebpmux3 \
    libopenexr23 \
    libopenslide0 \
    libheif1 \
    && apt autoremove -y && apt autoclean && apt clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=dev /usr/local/lib /usr/local/lib
COPY --from=dev /go/src/github.com/imgi/imgi/imgi /usr/local/bin/imgi
COPY tools/docker/imgi.conf /etc/imgi/imgi.conf

# Active vips libs
RUN ldconfig

CMD [ "imgi", "-c", "/etc/imgi/imgi.conf" ]

