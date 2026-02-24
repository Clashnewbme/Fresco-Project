# Welcome!

This guide will help you set up the Fresco browser.

# Step 1: Install Go (Use Shell to do so)

brew install go

# Step 2: Create the Project

mkdir FrescoBrowser
cd FrescoBrowser
go mod init frescobrowser
            
# Step 3: Install Webview

go get github.com/webview/webview_go
Step 4: Put this code in main.go

package main

import "github.com/webview/webview_go"

func main() {
    w := webview.New(true)
    defer w.Destroy()
    w.SetTitle("Fresco Browser")
    w.SetSize(800, 600, webview.HintNone)
    w.Navigate("https://google.com")
    w.Run()
}
           
# Step 5: Run

go run main.go











