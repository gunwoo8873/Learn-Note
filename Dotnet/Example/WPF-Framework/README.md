# Windows Presentation Foundation [WPF] Note

# Contents
[1. Getting Start](#1-getting-start)
[2. HTTP Server](#2-http-connect-go)

# 1. Getting Start GUI
* **Requirent**
  * [.NET SDK 9.0](https://dotnet.microsoft.com/ko-kr/download)
  * [Visual Studio Code](https://code.visualstudio.com/?wt.mc_id=developermscom)

* **Reference**
  * [Microsoft WPF : Window Presentation Foundation](https://learn.microsoft.com/ko-kr/dotnet/desktop/wpf/getting-started/?view=netframeworkdesktop-4.8)

* **CLI**
  * **Create Project**
    ```bash
    dotnet new wpf -n <Project name>
    ```
  
  * **Project Build**
    ```bash
    cd <Project name path>
    dotnet build
    ```

  * **Application Run**
    ```bash
    cd <Project name path>
    dotnet run
    ```

* **Path**
  ```md
  # WPF-Freamework/GUI
  /bin
  /obj
  App.xaml
  App.xaml.cs
  Assemblyinfo.cs
  MainWindow.xaml
  MainWindow.xaml.cs
  <Project name>.csproj # .csproj file name is CLI input value from created project name
  ```

# 2. HTTP Connect Go
* **Requirent**
  * [Go Download](https://go.dev/dl/)

* **Reference**
  * [HTTP Server Connection](https://pkg.go.dev/net/http)

* **CLI**
  * **Server Build**
    ```bash
    go build
    ```

* **Path**
  ```md
  # WPF-Freamework/Server
  go.mod
  go.sum
  main.go
  ```