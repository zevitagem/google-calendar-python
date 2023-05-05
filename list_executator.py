import datetime


class ListExecutator():
    def __init__(self, service):
        self.service = service

    def get_events(self):
        # Call the Calendar API
        now = datetime.datetime.utcnow().isoformat() + 'Z'  # 'Z' indicates UTC time
        
        print('Getting the upcoming 10 events')
        return self.service.events().list(
            calendarId='primary', timeMin=now,
            maxResults=10, singleEvents=True,
            orderBy='startTime'
        ).execute()

    def execute(self):
        events_result = self.get_events()
        events = events_result.get('items', [])

        if not events:
            return print('No upcoming events found.')

        # Prints the start and name of the next 10 events
        for event in events:
            start = event['start'].get('dateTime', event['start'].get('date'))
            print(start, event.get('summary', event['id']))
