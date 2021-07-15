from tkinter import *
import tkinter.messagebox
root = Tk()
root.title("nut.py")
root.geometry("500x500")




Checkbutton(text="python").pack(anchor=W)


# def show_choice():
#     ch = select.get()
#     if ch == 1: tkinter.messagebox.showinfo("Show", "Python Select")
#     elif ch == 2: tkinter.messagebox.showinfo("Show", "JS Select")
#     elif ch == 3: tkinter.messagebox.showinfo("Show","JAVA Select")
#     elif ch == 4: tkinter.messagebox.showinfo("Show","HTML Select")

# select = IntVar()
# Radiobutton(text="Python", variable=select, value=1, command=show_choice).grid(row=0,column=0)
# Radiobutton(text="JavaScrip",variable=select, value=2, command=show_choice).grid(row=0,column=1)
# Radiobutton(text="Java", variable=select, value=3, command=show_choice).grid(row=0,column=2)
# Radiobutton(text="HTML", variable=select, value=4, command=show_choice).grid(row=0,column=3)


root.mainloop()