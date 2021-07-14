from tkinter import *
import tkinter.messagebox
from tkinter.colorchooser import *
from tkinter.filedialog import *
root = Tk()
root.title("nut.py")
root.geometry("500x500")

def string():
    mess = us.get()
    la = Label(root, text=mess, fg="red").pack()

def exitProgram():
    comfirm = tkinter.messagebox.askquestion("Yes", "Exit ?")
    if comfirm == "yes": root.destroy()

def selectfile():
    fileopen = askopenfile()
    mylabel = Label(text=fileopen).pack()



us = StringVar()
txt = Entry(root, textvariable=us).pack()
Click = Button(root, text="Click", command=string).pack()
exit_root = Button(root, text="Exit", command=exitProgram).pack()


root.mainloop()