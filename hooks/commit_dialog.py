#!/usr/bin/env python3
import tkinter as tk
import urllib.request
import io
import sys

try:
    from PIL import Image, ImageTk
    PIL_AVAILABLE = True
except ImportError:
    PIL_AVAILABLE = False


def load_cat_image(root, url, size=(280, 280)):
    try:
        req = urllib.request.Request(url, headers={"User-Agent": "Mozilla/5.0"})
        with urllib.request.urlopen(req, timeout=5) as response:
            data = response.read()
        if PIL_AVAILABLE:
            img = Image.open(io.BytesIO(data)).resize(size)
            return ImageTk.PhotoImage(img)
        else:
            # Попытка без PIL (только PNG/GIF)
            photo = tk.PhotoImage(data=data)
            return photo
    except Exception:
        return None


result = {"value": 1}  # 1 = отмена по умолчанию


def on_yes():
    result["value"] = 0
    root.destroy()


def on_no():
    result["value"] = 1
    root.destroy()


root = tk.Tk()
root.title("Делать коммит?")
root.configure(bg="#1a1a2e")
root.resizable(False, False)

# Центрируем окно
root.update_idletasks()
w, h = 640, 480
x = (root.winfo_screenwidth() - w) // 2
y = (root.winfo_screenheight() - h) // 2
root.geometry(f"{w}x{h}+{x}+{y}")

# Заголовок
tk.Label(
    root, text="🐱 Делать коммит?", font=("Arial", 22, "bold"),
    bg="#1a1a2e", fg="white"
).pack(pady=(20, 5))

tk.Label(
    root, text="Выбери свою судьбу", font=("Arial", 13),
    bg="#1a1a2e", fg="#a0a0c0"
).pack(pady=(0, 15))

# Две картинки с подписями
frame_imgs = tk.Frame(root, bg="#1a1a2e")
frame_imgs.pack(pady=5)

# --- Левая картинка (ДА) ---
frame_yes = tk.Frame(frame_imgs, bg="#1a1a2e")
frame_yes.pack(side="left", padx=20)

yes_img = load_cat_image(root, "https://cataas.com/cat/cute", size=(240, 200))
if yes_img:
    lbl_yes_img = tk.Label(frame_yes, image=yes_img, bg="#1a1a2e")
    lbl_yes_img.image = yes_img
    lbl_yes_img.pack()
else:
    tk.Label(frame_yes, text="😺", font=("Arial", 80), bg="#1a1a2e").pack()

tk.Label(
    frame_yes, text="Да, всё огонь!", font=("Arial", 12, "bold"),
    bg="#1a1a2e", fg="#4CAF50"
).pack(pady=5)

tk.Button(
    frame_yes, text="✅  Пушим!", font=("Arial", 14, "bold"),
    bg="#4CAF50", fg="white", padx=25, pady=10,
    relief="flat", cursor="hand2", command=on_yes
).pack()

# --- Правая картинка (НЕТ) ---
frame_no = tk.Frame(frame_imgs, bg="#1a1a2e")
frame_no.pack(side="left", padx=20)

no_img = load_cat_image(root, "https://cataas.com/cat/angry", size=(240, 200))
if no_img:
    lbl_no_img = tk.Label(frame_no, image=no_img, bg="#1a1a2e")
    lbl_no_img.image = no_img
    lbl_no_img.pack()
else:
    tk.Label(frame_no, text="🙀", font=("Arial", 80), bg="#1a1a2e").pack()

tk.Label(
    frame_no, text="Нет, я облажался!", font=("Arial", 12, "bold"),
    bg="#1a1a2e", fg="#f44336"
).pack(pady=5)

tk.Button(
    frame_no, text="❌  Отмена", font=("Arial", 14, "bold"),
    bg="#f44336", fg="white", padx=25, pady=10,
    relief="flat", cursor="hand2", command=on_no
).pack()

root.mainloop()
sys.exit(result["value"])
