import npyscreen
import swagger_client
from swagger_client import AuthSuccess


class MaaSConfig:
    def __init__(self, host, verify_ssl):
        self.host = host
        self.verify_ssl = verify_ssl

class MaaS:
    def __init__(self, config: MaaSConfig):
        super().__init__()
        self._api_client = self.build_api_client(config)

    def build_api_client(self, config: MaaSConfig):
        # create an configuration for the general API client
        api_client_config = swagger_client.Configuration()
        api_client_config.host = config.host
        api_client_config.verify_ssl = config.verify_ssl

        # create an instance of the general API client
        return swagger_client.ApiClient(api_client_config)

    def build_auth_api_client(self):
        # create an instance of the API class
        return swagger_client.AuthApi(self._api_client)

    def build_items_client(self):
        return swagger_client.ItemsApi(self._api_client)


#class MatomatApp(npyscreen.NPSAppManaged):

    #def onStart(self):
        # self.addForm("MAIN", RecordListDisplay)
        # self.addForm("EDITRECORDFM", EditRecord)


if __name__ == '__main__':
    maas_cfg = MaaSConfig("https://localhost:8443/v0",
                          False)  # Do not do this for production use, only required to make it work with self signed certificates
    maas = MaaS(maas_cfg)
    auth_client = maas.build_auth_api_client()
    auth_response = auth_client.auth_login_post("admin", "admin")  # type: AuthSuccess
    token = auth_response.token
    items_client = maas.build_items_client()
    items_client.items_post("mate", 100, token=token)
    items_client.items
    #myApp = MatomatApp()
    #myApp.run()
