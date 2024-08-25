import requests
import re
import pandas as pd
from bs4 import BeautifulSoup

def make_request(payload):
    """
    Function to send a GET request to the Google Search API and handle potential errors.
    :param payload: Dictionary containing the API request parameters
    :return: JSON response from the API
    """
    response = requests.get('https://www.googleapis.com/customsearch/v1', params=payload)
    if response.status_code != 200:
        raise Exception('Request failed')
    return response.json()

def build_payload(query, start=1, num=10, **params):
    """
    Function to build the payload for the Google Search API request.
    
    :param query: Search term
    :param start: The index of the first result to return
    :param num: The number of results to return (default is 10)
    :param date_restrict: Restricts results based on recency (default is one month 'm1')
    :param params: Additional parameters for the API request
    
    :return: Dictionary containing the API request parameters
    """
    payload = {
        'key': API_KEY,
        'q': query,
        'cx': SEARCH_ENGINE_ID,
        'start': start,
        'num': num,
    }
    payload.update(params)
    return payload


def find_email_by_name_and_institution(url):
    
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'html.parser')

    # Find and parse email addresses
    emails = []
    for word in soup.stripped_strings:
        if "@" in word and "." in word:
            emails.append(word)
    
    return emails

def main(query, result_total=10):
    """
    Main function to execute the script and print all URLs from the search results.
    """
    items = []
    reminder = result_total % 10
    if reminder > 0:
        pages = (result_total // 10) + 1
    else:
        pages = result_total // 10
    
    for i in range(pages):
        if pages == i + 1 and reminder > 0:
            payload = build_payload(query, start=(i+1)*10, num=reminder)
        else:
            payload = build_payload(query, start=(i+1)*10)
        response = make_request(payload)
        items.extend(response['items'])
    
    urls = [item['link'] for item in items]  # Extract URLs from the response
    
    return urls


if __name__ == "__main__":
    API_KEY= 'AIzaSyAa3v8ulaMd6MXQ1oCJDzNCG4pHV6Ms8OU'
    SEARCH_ENGINE_ID = '227c94475aca5432c'
    search_query = "Chris Mcdonald University of Western Australia"
    total_results = 20

    urls = main(search_query, total_results)
    print("URLs found:")
    for url in urls:
        if "pdf" in url:
            continue
        # print(url)

        try:
            if len(find_email_by_name_and_institution(url)[0]) > 100:
                continue
            else:
                # print(len(find_email_by_name_and_institution(url)))
                print(url)
                print(find_email_by_name_and_institution(url))
        except:
            continue
