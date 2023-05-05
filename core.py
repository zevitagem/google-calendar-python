from __future__ import print_function

import datetime
import os.path
import sys

from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError

from list_executator import ListExecutator
from insert_executator import InsertExecutator


def auth():
    # If modifying these scopes, delete the file token.json.
    SCOPES = ['https://www.googleapis.com/auth/calendar']

    """Shows basic usage of the Google Calendar API.
    Prints the start and name of the next 10 events on the user's calendar.
    """
    creds = None
    token_file = './token.json'
    credentials_file = './credentials.json'
    
    # The file token.json stores the user's access and refresh tokens, and is
    # created automatically when the authorization flow completes for the first
    # time.
    if os.path.exists(token_file):
        creds = Credentials.from_authorized_user_file(token_file, SCOPES)

    # If there are no (valid) credentials available, let the user log in.
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(
                credentials_file, SCOPES)
            creds = flow.run_local_server(port=0)

        # Save the credentials for the next run
        with open(token_file, 'w') as token:
            token.write(creds.to_json())

    return creds


def main():
    try:
        creds = auth()
        service = build('calendar', 'v3', credentials=creds)

        _, action = tuple(sys.argv)
        if (action == 'list'):
            executator = ListExecutator(service)

        if (action == 'insert'):
            executator = InsertExecutator(service)

        executator.execute()

    except HttpError as error:
        print('An error occurred: %s' % error)


if __name__ == '__main__':
    main()
