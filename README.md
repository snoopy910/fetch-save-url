# fetch-save-url

## Problem

Implementation of a command line program that can fetch web pages and saves them to disk for later retrieval and browsing.

### Section 1

For example, if we invoked your program like this: `./fetch https://www.google.com` then in our current directory we should have a file containing the contents of `www.google.com`. (i.e. `/home/myusername/www.google.com.html`).

We should be able to specify as many urls as we want:

```
$> ./fetch https://www.google.com https://autify.com <...>
$> ls
autify.com.html www.google.com.html
```

If the program runs into any errors while downloading the html it should print the error to the console.

### Section 2

Record metadata about what was fetched:

- What was date and time of last fetch
- How many links are on the page
- How many images are on the page

Modify the script to print this metadata.

For example (it can work differently if you like)

```
$> ./fetch --metadata https://www.google.com
site: www.google.com
num_links: 35
images: 3
last_fetch: Tue Mar 16 2021 15:46 UTC
```

## Prerequisities

You have to install docker before building and running

## How to build and run

- Follow the command to build and run the docker image

```
docker build fetch .
docker run fetch ./fetch {target_url}
docker run fetch ./fetch --metadata {target_url}
```
