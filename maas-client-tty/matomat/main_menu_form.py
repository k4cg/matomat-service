import npyscreen

# ----------
# main menu
# ----------
# <<<Mate-Logo KLEIN>>>
# <<Welcome>>
# <<Your credits are:>>
#
# (1-4.) most used items
# 5. more items ...
# 6. credits
# 7. stats
# 8. account
# if admin
#     9. user management
#     10. item management
#     11. logout
# else
#     9. logout
class MainMenuForm(npyscreen.Form):

    def __init__(self, name=None, parentApp=None, framed=None, help=None, color='FORMDEFAULT', widget_list=None,
                 cycle_widgets=False, *args, **keywords):
        super().__init__(name, parentApp, framed, help, color, widget_list, cycle_widgets, *args, **keywords)

    def create(self):
        self.action = F.add(npyscreen.TitleSelectOne, max_height=4, value=[1, ], name="What's next?",
                   values=["4", "Option2", "Option3"], scroll_exit=True)
