# zippy

Born out of anger and frustration. See the [wiki](https://github.com/tugberkugurlu/zippy/wiki) for current nonsense.

## What I want to do

 - Package up a zip/tar/{insert-archive-format} file and send it to an HTTP endpoint with a version information: `zippy push ./my-package-name.zip --version 1.2.3-beta.56423`
 - Be able to say `zippy get my-package-name --out ./tools/` to get the latest.

## Building

In order to build the application locally, you need the following tools:

 - **ruby**: `sudo apt-get install ruby`
 - **rake**: `gem install rake`
 - **dnx**: `dnvm install 1.0.0-beta6-12232 -r mono -u`
