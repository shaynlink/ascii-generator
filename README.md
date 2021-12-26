# ASCII API

Generate ASCII art from image.

You can try this API here:
[soon]()

# Installation

## Prerequisites
 - [Golang 1.17](https://go.dev)
 - [Git](https://git-scm.com) (for clone project)

```bash
git clone https://github.com/shaynlink/ascii-api.git

cd ascii-api

go run *.go
```

I recommend to use [Postman](https://www.getpostman.com/). for try API in local

# Enpoints

| Method | Path | Description |
| ------ | ---- | ----------- |
| GET | / | Generate ASCII art from image. |
| * | * | Generate ASCII art from image. |

# Parameters

## (GET) /
Parameters **form-data**
 - **image**: Image file.
 - **width**: Width of the generated ASCII art.
 - **invert**: Invert the image.

Return **text ACII art**.

# Example code client

### Curl

### NodeJS (axios)
