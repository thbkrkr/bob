#!/usr/bin/env python

from flask import Flask
app = Flask(__name__)

@app.route("/")
def root():
  return app.send_static_file('index.html')

@app.route("/api/forward")
def forward():
  return "Forward"

@app.route("/api/back")
def back():
  return "back"

@app.route("/api/left")
def left():
  return "Left"

@app.route("/api/right")
def right():
  return "Right"

if __name__ == "__main__":
  app.run(host='0.0.0.0')