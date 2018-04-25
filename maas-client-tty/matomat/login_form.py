import npyscreen

# ----------
# login
# ----------
# <<<Mate-Logo GROSS>>>
class LoginForm(npyscreen.Form):
    def create(self):
        self.login = self.add(npyscreen.TitleText, name='Login')
