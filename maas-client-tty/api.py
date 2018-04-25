from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from swagger_client.models.auth_success import AuthSuccess
from pprint import pprint

# create an configuration for the general API client
api_client_config = swagger_client.Configuration()
api_client_config.host = "https://localhost:8443/v0"
api_client_config.verify_ssl = False #Do not do this for production use, only required to make it work with self signed certificates

# create an instance of the general API client
api_client = swagger_client.ApiClient(api_client_config)

# create an instance of the API class
api_instance = swagger_client.AuthApi(api_client)

try:
    # DUMMY TEST CODE
    # Logs a user in and returns an JWT token for authentication
    username = "admin"
    password = "admin"
    validitySeconds = 7200
    result = api_instance.auth_login_post(username, password, validityseconds=validitySeconds) # type: AuthSuccess
    pprint(result.token)
    pprint(result)
except ApiException as e:
    print("Exception when calling AuthApi->auth_login_post: %s\n" % e)