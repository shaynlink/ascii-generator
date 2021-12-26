# ASCII API

Generate ASCII art from image.

You can try this API here:
[ascii.projects.shaynlink.dev](http://ascii.projects.shaynlink.dev) or if you prefer using HTTPS [ascii-generator.herokuapp.com](https://ascii-generator.herokuapp.com)

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
| POST | / | Generate ASCII art from image. |
| * | * | Generate ASCII art from image. |

# Parameters

## (POST) /
Parameters **form-data**
 - **image**: Image file.
 - **width**: Width of the generated ASCII art.
 - **invert**: Invert the image.

Return **text ACII art**.

# Example code client

### Curl

```bash
curl --location --request POST 'http://ascii.projects.shaynlink.dev' \
--form 'image=@"./test.png"' \
--form 'width="80"' \
--form 'invert="true"'
```

### NodeJS (axios)

```javascript
const axios = require('axios');
const FormData = require('form-data');
const fs = require('fs');

const data = new FormData();

data.append('image', fs.createReadStream('./test.png'));
data.append('width', 80);
data.append('invert', true);

axios({
  method: 'post',
  url: 'http://ascii.projects.shaynlink.dev',
  headers: data.getHeaders()
  data
})
    .then(({data}) => {
        console.log(data);
    });
```

# Credits
All of the [contributors](https://github.com/shaynlink/ascii-generator/graphs/contributors) that have helped with implementing various features, adding themes, fixing bugs, and more.

# Support
Thanks you for support my project and me, using ascii-generator.

You can also support me by [donate](https://paypal.me/shaynl4nk), [patron](https://www.patreon.com/shaynlink) or [buy me a coffee](https://www.buymeacoffee.com/shaynlink).