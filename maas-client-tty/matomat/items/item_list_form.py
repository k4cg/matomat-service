import npyscreen
from matomat.form_ids import FormIDs

class ItemList(npyscreen.MultiLineAction):
    def __init__(self, *args, **keywords):
        super().__init__(*args, **keywords)
        self.add_handlers({
            "^A": self.when_add_record,
            "^D": self.when_delete_record
        })

    def display_value(self, vl):
        return "%s, %s" % (vl[1], vl[2])

    def actionHighlighted(self, act_on_this, keypress):
        self.parent.parentApp.getForm(FormIDs.ITEM_EDIT_FRM).value =act_on_this[0]
        self.parent.parentApp.switchForm(FormIDs.ITEM_EDIT_FRM)

    def when_add_record(self, *args, **keywords):
        self.parent.parentApp.getForm(FormIDs.ITEM_CREATE_FRM).value = None
        self.parent.parentApp.switchForm(FormIDs.ITEM_CREATE_FRM)

    def when_delete_record(self, *args, **keywords):
        self.parent.parentApp.myDatabase.delete_record(self.values[self.cursor_line][0])
        self.parent.update_list()