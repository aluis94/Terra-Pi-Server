from appConfig import *

from twilio.rest import Client
client = Client(account_sid, auth_token)
message = client.api.messages.create(
    to="+13109772732",
    from_="+18665254691",
    body="This is a test SMS message"

)
