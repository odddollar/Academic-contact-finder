

<img src="../images/Big_Header.png" style="margin: auto;">





<h1 style="text-align: center; font-size: 70px; color: #6576A8;">User Guide</h1>































## Table of Contents

[TOC]

## 1. Introduction

Welcome to the **Academic Contact Finder**. This application is a user friendly way for finding email addresses of academic researchers, doctors, or professors. 



## 2. System Requirements 

- **Operating System**: Windows 10 or later, macOS 10.15 or later, Linux
- **Memory**: Minimum 4 GB RAM
- **Storage**: Minimum 100 MB of available space
- **Internet Connection**: Required for API requests and web scraping
- **API Key**: A valid Scopus API key **OR** A valid Google API key
  - If you provide both you be able to attain the most accurate results




## 3. Installation

### Download the Software

- visit the GitHub repository and download the latest version

### Launch the Software 

* After installataion, simply running the **.exe** file will open the application 



## 4. Initial Setup 

Upon the first time using the application, you will be prompted to enter your **Scopus API key** and your **Google API key**.

1. **Enter the API Key**:
   - Copy the API key you have been provided through Scopus and the one you have attained through Google. 
   - Paste the keys into the respective fields when prompted.

2. **Save the API Key**:
   - Click the "Save" button to store your API keys. 



## 5. Using the App 

### Entering Researcher Details 

1. **Enter Name in Fields Provided**:
   - Provide first and last name for the most accurate results.
   - However, if only a last name is entered, the search will still be performed.
2. **Enter Institution Details**:
   - Provide the name for the researcher's institution if that information is known.

### Performing the Search

1. **Initiate Search**:
   - After entering the necessary details press the "Search" button.
2. **Processing**:
   - The app will perform API requests from "Scopus" first.
   - If limited results is found through this method it will automatically perform web scraping.

### Viewing Results 

1. **Review the Results**:
   - Once search is complete, the results will be displayed at the bottom of the screen.

2. **Order of Results**:
   - You may sort results in order based on publication dates. 

## 6. Emailing Results

1. **Setting Default Email Address**:
   - Press the button "Change default email address" to change your email which is currently entered.
2. **Email All Results**:
   - By pressing "Email All" at the bottom of the screen, you will then be allowed to select "Outlook". 
   - Once selected, you will be able to send all results that the app has found to your email address

3. **Send Individual Results**:
   - If you only wish to email yourself one of the results, simply press the "Send" button inside the result you wish to choose

## 7. Managing the Scopus API Key



## 8. Troubleshooting



## 9. FAQs

- **What is an API key?**

    An API key is a unique identifier used to authenticate a user, developer, or program to an API (Application Programming Interface). It allows our application to communicate with external services, namely Scopus and Google, and access data based on the permissions granted by the key.

- **Why do I need an API key?**

    You will need an API key to access the data from Scopus and Google that provides the contact details of researchers. The API keys verify that our application has the proper authorisation to use these services, ensuring security and monitoring usage.

- **Why do I need two different API Keys?**

    Different services require their own API keys for access. The Scopus API key allows our application to search for researcher details within the Scopus database, while the Google API key is used to find publicly available information through Google Search. Each service has its own API and needs its own key for interaction.

    For best results, both API keys should be provided.

- **Should I be sharing my Google API key with others?**

    No, you should not share your Google API key with others. Sharing your API key can expose your account to security risks. For this reason, it is important to keep your API key private.

    If others wish to use the ‘Academic Research Finder,’ they will need to generate their own Google API key by following the instructions provided above. 

- **How up to date is the information provided?**

    While our application aims to provide the most up-to-date information for each researcher, it will only return results from Scopus or Google. This does not guarantee that all information is current.

- **How accurate is the information provided?**

    While our application aims to provide the most accurate information for each researcher, it can only return results found from Scopus or Google. 

    There may be instances where the associated email appears incorrect. Since we provide the source from which the details were obtained, this can be manually verified.

- **How does the application find contact details?**

    The application uses a combination of API calls and web scraping to gather publicly available information from Scopus, an academic database, and websites online.

- **Is the data I input stored or shared?**

    No, the data you input is not stored or shared. The application processes your input only temporarily to search for contact details. Once the search is completed and the results are delivered, the input data is not retained or shared.

- **Is there a limit to how many searches I can perform?**

    While there is no limit to how many searches can be performed by Scopus API key, the Google API key is limited to 100 search queries per day for free. 

- **Is there any way to copy the results directly to my clipboard?**

    Yes, clicking the clipboard icon on the right-hand side of the results will copy a formatted version of the results to your clipboard.

- **Can I search for researchers worldwide, or are there geographic restrictions?**

    The application can search for researchers globally, provided that their contact information is publicly available by the services used. 

- **Why can’t I find the researchers' contact details?**

    If the researcher's contact details are not publicly accessible or listed in Scopus, the application will not be able to retrieve the information. 

    Otherwise, ensure that the information you have entered in the search fields is spelt correctly and uses the researchers full name as opposed to nicknames. For institutions, use the full name of the institution rather than abbreviations. For example, enter the ‘University of Western Australia’ instead of ‘UWA.’

- **Which Wi-Fi network should I be connected to in order to use this?**

    To access results from Scopus, you need to be connected to a Wi-Fi network that provides access to Scopus, such as your office network. If you are not connected to such a network, the ‘Academic Contact Finder’ will return results from Google Search instead.