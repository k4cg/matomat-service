import npyscreen

from matomat.login_form import LoginForm


def loginFunc(*args):
    F = LoginForm(name="Please provide your credentials to log into Matomat")
    F.edit()
    return "Logged in " + F.login.value


if __name__ == '__main__':
    print(npyscreen.wrapper_basic(loginFunc))
