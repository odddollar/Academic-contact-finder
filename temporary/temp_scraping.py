from selenium import webdriver
from bs4 import BeautifulSoup
from time import sleep

# Create a new instance of the web driver (Chrome in this case)
driver = webdriver.Firefox()

# The URL you want to visit
url = "https://www.scopus.com/record/display.uri?eid=2-s2.0-85079320615&origin=resultslist&sort=plf-f&src=s&st1=McDonald&st2=Chris&nlo=1&nlr=20&nls=count-f&sid=2a239e24d72b73730a45aca45c5e6c37&sot=anl&sdt=aut&sl=39&s=AU-ID%28%22McDonald%2c+Chris+S.%22+57169566400%29&relpos=2&citeCnt=133&searchTerm="

# Use the driver to visit the webpage
driver.get(url)

# Make a bit more reliable
sleep(1)

# Parse the HTML with BeautifulSoup
soup = BeautifulSoup(driver.page_source, "html.parser")

lis = soup.find_all("li")

# Find all anchor tags
for li in lis:
    a = li.find_all("a", href=True)

    # Iterate through mailto links
    for i in a:
        if "mailto" in i["href"]:
            # Find span within button to get name
            name = li.find("button").find("span")
            print(name.text + ": " + i["href"][7:])

driver.quit()
