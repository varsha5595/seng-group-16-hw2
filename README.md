# seng-group-16-hw2
CSC510 Software Engineering Homework-2

## Getting Started

These instructions will help you in setting up the project and running the game on your local machine for development and testing. 

## Language Installation

### Ruby

#### MacOS
```
brew install ruby
ruby -v
```

#### Linux
```
sudo apt-get install ruby-full
ruby -v
```

#### Windows
Download the installation file from https://www.ruby-lang.org/en/downloads/
Add the following lines to your c:\autoexec.bat: set PATH = "D:\(ruby install directory)\bin;%PATH%"

```
ruby -v
```

### Golang

#### MacOS
```
brew install golang
```

#### Linux

Download the Linux distribution from [Go’s official download page](https://golang.org/dl/) and extract it into /usr/local directory.

```
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```
Next, add the /usr/local/go/bin directory to your PATH environment variable. You can do this by adding the following line to your ~/.bash_profile file -

```
export PATH=$PATH:/usr/local/go/bin
```

#### Windows

Download the Windows MSI installer file from [Go’s official download page](https://golang.org/dl/). Open the installer and follow the on-screen instructions to install Go in your windows system. By default, the installer installs Go in ```C:\Go```

## Setting GOPATH

#### Unix Systems (Linux and macOS)

For setting GOPATH in bash shells, add the following line to the ~/.bash_profile file -

```
export GOPATH=$HOME/go
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

If you use Zsh shell, then you need to add the above line to ~/.zshrc file.

#### Windows
Copy the path of the ```seng-group-16-hw2``` project folder. Now set the ```GOPATH``` environment variable using the following instrustions:

1. Right click on Start → click Control Panel → Select System and Security → click on System
2. From the menu on the left, select the Advanced systems settings
3. Click the Environment Variables button at the bottom
4. Click New from the User variables section
5. Type GOPATH into the Variable name field
6. Paste the ```seng-group-16-hw2```  folder path into the Variable value field
7. Click OK

Note that, GOPATH must be different than the path of your Go installation.


### Rust

#### MacOS 
```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

rustc --version
cargo --version
```

#### Linux
```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

rustc --version
cargo --version
```

#### Windows

Setup Rust with the following executable file https://static.rust-lang.org/rustup/dist/i686-pc-windows-msvc/rustup-init.exe

Once you open it, you’ll be met with the installer options. Type "Y" and let the installer do a set of downloads

Add the following to PATH in environment variables
```
%USERPROFILE%\.cargo\bin
```
Check the installation with the following
```
rustc --version
cargo --version
```
## Execution

### Ruby

```
ruby ruby.rb
```

### Golang
```
cd go/src
go run golang.go
```

### Rust
```
rustc rust.rs
./rust
```

## Authors

* Adithya Raghu Ganesh
* Meghana Ravindra Vasist
* Shivaprakash Balasubramanian
* Surbhi Jha
* Varsha Anantha Ramu Sharma

## License

This project is licensed under the MIT License.