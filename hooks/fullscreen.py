#!/usr/bin/env python3
import tkinter as tk
import subprocess
import threading
import os
import sys

def play_sound():

    subprocess.run(["osascript", "-e", "set volume output volume 100"], capture_output=True)
    subprocess.run(["afplay", "/System/Library/Sounds/Sosumi.aiff"])
    subprocess.run(["afplay", "/System/Library/Sounds/Sosumi.aiff"])


script_dir = os.path.dirname(os.path.abspath(__file__))
project_root = os.path.join(script_dir, "..", "..")
image_path = os.path.join(project_root, "download.jpg")

root = tk.Tk()
root.attributes("-fullscreen", True)
root.configure(bg="black")
root.bind("<Escape>", lambda e: root.destroy())
root.after(3000, root.destroy)

try:
    from PIL import Image, ImageTk
    img = Image.open(image_path)
    screen_w = root.winfo_screenwidth()
    screen_h = root.winfo_screenheight()
    img = img.resize((screen_w, screen_h))
    photo = ImageTk.PhotoImage(img)
    tk.Label(root, image=photo, bg="black").pack(expand=True)
except Exception:
    tk.Label(root, text="🐰 ЧЕРЕМША 🐰", font=("Arial", 100, "bold"),
             bg="black", fg="white").pack(expand=True)


threading.Thread(target=play_sound, daemon=True).start()

root.mainloop()
