from twilio.rest import Client
from appConfig import *
#account_sid = "ACb221ea80db1b430855f4ffa09136b5a5" 
#auth_token = "e340918a56c17920c5c10049cb3a2854"

client = Client(account_sid, auth_token)
message = client.api.messages.create(
    to="+13109772732",
    from_="+18665254691",
    body="This is a test mesage"

)