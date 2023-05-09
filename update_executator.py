class UpdateExecutator():
    def __init__(self, service):
        self.service = service
        self._id = None

    def configure(self, config: dict):
        self._id = config.get('id', None)

    def execute(self):
        events_service = self.service.events()

        id = self._id #'827gj35869mlimobovmajc1m94'
        event = events_service.get(calendarId='primary', eventId=id).execute()

        event['summary'] = 'Appointment at Somewhere - Update'

        updated_event = events_service.update(calendarId='primary', eventId=event['id'], body=event).execute()
        
        # Print the updated date.
        print(updated_event['updated'])
