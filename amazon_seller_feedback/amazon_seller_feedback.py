import requests
from urllib.parse import urlparse, parse_qs
import sys

amazon_url = input("Please input Amazon Seller URL: ")
try:
	parsed_url = urlparse(amazon_url)
	marketplaceID = parse_qs(parsed_url.query)['marketplaceID'][0]
	seller = parse_qs(parsed_url.query)['seller'][0]
except:
	print("Invalid URL\nPlease make sure this is a Amazon seller page and try again.\n\n")
	sys.exit(1)

print("starting work...")
next_page = True #Starts true so while loop execuxes
page_number = 1
form_data = { 
	"seller":str(seller),
	"marketplaceID":str(marketplaceID),
	"pageNumber":str(page_number)
}

while next_page:
	r = requests.post("https://www.amazon.com/sp/ajax/feedback",data=form_data)
	resp = r.json()
	next_page = resp['hasNextPage']
	page_number += 1
	form_data = { 
		"seller":str(seller),
		"marketplaceID":str(marketplaceID),
		"pageNumber":str(page_number)
	}


	for review in resp['details']:
		if review['rating'] < 3:
			print("%s: %s\n%s\n\n" % (review['ratingData']['date'], review['rating'],review['ratingData']['text']['expandedText']))


print("Done.\n\n")
