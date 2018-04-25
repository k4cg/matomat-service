import npyscreen
import swagger_client


def buildApiClient():
    # create an configuration for the general API client
    api_client_config = swagger_client.Configuration()
    api_client_config.host = "https://localhost:8443/v0"
    api_client_config.verify_ssl = False  # Do not do this for production use, only required to make it work with self signed certificates

    # create an instance of the general API client
    return swagger_client.ApiClient(api_client_config)

class MatomatApp(npyscreen.NPSAppManaged):
    def onStart(self):
        self.apiClient = buildApiClient()
        self.addForm("MAIN", RecordListDisplay)
        self.addForm("EDITRECORDFM", EditRecord)

if __name__ == '__main__':
    myApp = AddressBookApplication()
    myApp.run()