# Welcome!
This guide will help you set up the Fresco Project.

# Step 1: Install Go (Use Command Prompt to do so)
Go to https://go.dev/dl/ and download the Windows installer (.msi).
Run the installer and follow the setup.

# Step 2: Create the Project
mkdir FrescoProject
cd FrescoProject
go mod init frescoproject

# Step 3: Install Webview
go get github.com/webview/webview_go

# Step 4: Open all the files
Download and open main.go inside the FrescoProject folder.
Make sure WebView2 runtime is installed: https://developer.microsoft.com/en-us/microsoft-edge/webview2/

# Step 5: Run
go run main.go

# Step 6 (If you want): Build EXE
go build -o Fresco.exe
Double-click Fresco.exe to run your browser.
