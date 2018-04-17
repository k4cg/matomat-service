from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.AuthApi()

try:
    # DUMMY TEST CODE
    # Logs a user in and returns an JWT token for authentication
    username = "admin"
    password = "admin"
    validitySeconds = 7200
    result = api_instance.auth_login_post(username, password, validityseconds=validitySeconds)
    pprint(result)
except ApiException as e:
    print("Exception when calling AuthApi->auth_login_post: %s\n" % e)