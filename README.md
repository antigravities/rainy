# Rainy

A (WIP) file-sharing service, similar in spirit to Pomf.

## Install

### Standalone

Install Go, Node.js, and Git. Then:

```bash
git clone https://github.com/antigravities/rainy.git
cd rainy

# Build the React-based frontend
cd www && npm install && npm run build && cd ..

# Build the app itself (embeds frontend)
go build
```

Finally, create a `.env` file with your favorite text editor. The variables are:

```bash
RAINY_UPLOADER=filesystem
# The uploader to use. Only `filesystem` is available at this time, but support
# for other storage mechanisms (i.e. S3) is planned.

RAINY_UPLOADER_PATH=./uploads
# The place to store uploads. It will be created if it doesn't exist.

RAINY_MAX_FILE_SIZE=10000000
# The maximum file size Rainy will accept.

RAINY_SHOULD_SERVE=1
# Whether Rainy should serve files itself. When this is enabled, files will be
# served from http://rainy-host/f.

RAINY_INSTANCE_NAME=Rainy
# The name of this instance of Rainy.

RAINY_BLACKLISTED_EXTENSIONS=exe,jar,ipa,apk
# File extensions Rainy will not accept.

RAINY_UPLOAD_PASSWORD=Rainy
# Require a password to upload files. The user will be prompted to enter this
# password when uploading. Leaving this empty will allow anyone to upload
# files.

RAINY_PUBLIC_UPLOAD_URL=https://eeti.me/f
# The URL Rainy should return for uploaded files.
```

You're all set! Run

```
./rainy
```

### Docker

Create a `.env` file as illustrated above. Then pull and run
`antigravities/rainy:latest`. A sample docker-compose.yml file is provided.

## Hacking

The frontend is a React app created with `create-react-app`, and all of the
normal `react-scripts` procedures are still intact.

Rainy will automatically replace the bundled files if it can find a replacement
locally, (i.e. if a file called `html/boomer.html` exists, Rainy will use that
instead of the bundled file).

## License

```
Rainy, a file sharing service
Copyright (C) 2021 antigravities

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```