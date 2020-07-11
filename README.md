# imgi

**imgi** is an image procressing service.

## Installation

```
go get -u github.com/imgi/imgi/cmd/imgi
```

## Usage

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


## License

MIT - Copyright (c) 2020 Chen Guohui

