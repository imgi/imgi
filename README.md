# imgi

[![Travis Build Status](https://travis-ci.com/imgi/imgi.svg?branch=master)](https://travis-ci.com/imgi/imgi)

**imgi** is an image procressing service.

## Installation

You can install imgi from [docker](#docker) or [build from source](#build-from-source).

### Build from source

#### Prerequisite

* GO 1.14+
* libvips 8.9+

#### Build and install

```
go get -u github.com/imgi/imgi/cmd/imgi
```

## Usage

### Docker

Start imgi container with default configuration:
```
docker run -d -p 6969:6969 -v /path/to/images:/images imgi/imgi
```

### Example

Start imgi server with a custom configuration file:
```
imgi -c imgi.conf
```

Resize an image:
```
http://localhost:6969/images/image.jpg?imgi=resize:w_640,h_480,m_fit
```

### API

API is access by providing an http query string `imgi` with the following format:
```
imgi=operation0:option0_value0,option1_value1|operation1:option0_value0
```

## Operations

### Resize `resize`

| Option | Type | Value | Description |
| ------ | ---- | ----- | ----------- |
| m | string | `fit, fill, flex` default: `fit` | content mode: <br/> `fit`: aspect ratio preserved, width and height attain <br/> `fill`: aspect ratio ignored, width and height attain <br/> `flex`: aspect ratio preserved, width or height attain |
| w | int | `[1, 10000]` | output image width |
| h | int | `[1, 10000]` | output image height |


### Crop `crop`

| Option | Type | Value | Description |
| ------ | ---- | ----- | ----------- |
| w | int | `[1, width - x]` | output image width |
| h | int | `[1, height - y]`| output image height |
| x | int | `[0, width]` | x-axis along image width |
| y | int | `[0, height]` | y-axis along image height |

## License

MIT - Copyright (c) 2020 Chen Guohui

