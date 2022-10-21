# Amazon Seller feedback

This script was created because seeing poor reviews on Amazon's seller pages is very difficult, time consuming, and tedious.

This script takes an Amazon seller URL and prints all the bad (1 or 2 star ratings) reviews of that seller including the date, rating and full review text.

# Installation
1. [Install pyenv](https://github.com/pyenv/pyenv#installation)
1. Install the latest version of Python 3: `pyenv install 3:latest`
1. Clone this repo and navigate to this folder
1. Use the latest version of Python 3 in this folder. e.g. `pyenv local 3.10.7`
1. Install dependencies: `pip install -r requirements.txt`

# Usage

`python amazon_seller_feedback`

Example:
```
$ python amazon_seller_feedback.py
Please input Amazon Seller URL: https://www.amazon.com/sp?marketplaceID=ATVPDKIKX0DER&seller=A3DJTZRS1C5A82&isAmazonFulfilled=1&ref_=dp_merchant_link&asin=B07WPTG7NX
"starting work..."
[date]: [rating]
[review text]

...

Done.
```
