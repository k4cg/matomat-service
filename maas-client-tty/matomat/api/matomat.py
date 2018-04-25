import swagger_client
from swagger_client.models import AuthSuccess
from swagger_client.models import Error


class Matomat():
    def __init__(self, api_client: swagger_client.ApiClient):
        super().__init__()
        self.api_client = api_client
        self.auth_success = None  # type: AuthSuccess

    def login(self, username, password, token_validity_seconds):
        api_instance = swagger_client.AuthApi(self.api_client)
        try:
            result = api_instance.auth_login_post(username, password, validityseconds=token_validity_seconds)
            if isinstance(result, Error):
                raise RuntimeError(Error.message)
        except Exception as e:
            raise e
        self.auth_success = result

    def listItems(self):
        self.
