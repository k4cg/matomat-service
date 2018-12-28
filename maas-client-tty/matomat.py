import npyscreen
import swagger_client
from swagger_client import AuthSuccess


class MaaSConfig:
    def __init__(self, host, verify_ssl):
        self.host = host
        self.verify_ssl = verify_ssl

class MaaSApiClientBuilder:
    def __init__(self, config: MaaSConfig):
        super().__init__()
        self._maas_config = config

    def build_auth_api_client(self):
        # create an instance of the API class
        return swagger_client.AuthApi(swagger_client.ApiClient(self.build_config()))

    def build_items_client(self, token):
        return swagger_client.ItemsApi(swagger_client.ApiClient(self.build_config_with_token(token)))

    def build_config(self):
        # create an configuration for the general API client
        api_client_config = swagger_client.Configuration()
        api_client_config.host = self._maas_config.host
        api_client_config.verify_ssl = self._maas_config.verify_ssl

        return api_client_config

    def build_config_with_token(self, token):
        api_client_config = swagger_client.Configuration()
        api_client_config.host = self._maas_config.host
        api_client_config.verify_ssl = self._maas_config.verify_ssl
        api_client_config.api_key = {
            'Authorization': token
        }
        api_client_config.api_key_prefix = {
            'Authorization': 'Bearer'
        }

        return api_client_config

#class MatomatApp(npyscreen.NPSAppManaged):

    #def onStart(self):
        # self.addForm("MAIN", RecordListDisplay)
        # self.addForm("EDITRECORDFM", EditRecord)


if __name__ == '__main__':
    maas_cfg = MaaSConfig("https://localhost:8443/v0",
                          False)  # Do not do this for production use, only required to make it work with self signed certificates
    maas_builder = MaaSApiClientBuilder(maas_cfg)
    auth_client = maas_builder.build_auth_api_client()
    auth_response = auth_client.auth_login_post("admin", "admin")  # type: AuthSuccess
    token = auth_response.token
    items_client = maas_builder.build_items_client(token)
    response = items_client.items_post("mate", 100)
    #myApp = MatomatApp()
    #myApp.run()
    print(response)
