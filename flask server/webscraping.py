import requests 
# import beautifulsoup

#API Key: AIzaSyAa3v8ulaMd6MXQ1oCJDzNCG4pHV6Ms8OU

def queryGoogleAPI(parameters):
    # The API endpoint
    url = "https://www.googleapis.com/customsearch/v1?key=AIzaSyAa3v8ulaMd6MXQ1oCJDzNCG4pHV6Ms8OU&cx=017576662512468239146:omuauf_lfve&q="+f'{parameters}'

    # A GET request to the API
    response = requests.get(url)

    return response


def parsing(htmltext):
    response2 = requests.get(url2)
    # Print the response
    # print(response.json())
    html = response2.text

    htmlsplit = html.split(" ")
    for i in htmlsplit:
        if i == "<a":
            print(i)



def main():
    url2 = "https://www.linkedin.com/in/chrismcdonald238/overlay/contact-info/"
    parsing(url2)



if __name__ == "__main__":
    main()    