#!/bin/bash

# In this project, protofiles are left inside the repository, they are only regenerated if they are changed

get_package_manager() {
    if command -v apt &> /dev/null; then
        echo "apt"
    elif command -v pacman &> /dev/null; then
        echo "pacman"
    elif command -v yum &> /dev/null; then
        echo "yum"
    elif command -v dnf &> /dev/null; then
        echo "dnf"
    else
        echo "unsupported"
    fi
}

install_protoc() {
    local package_manager="$1"
    
    case "$package_manager" in
        apt)
            sudo apt update
            sudo apt install -y protobuf-compiler
            ;;
        pacman)
            sudo pacman -Syu protobuf
            ;;
        yum|dnf)
            sudo $package_manager install protobuf
            ;;
        *)
            echo "Unsupported package manager. Please install protoc manually."
            exit
            ;;
    esac
}

# Check if protoc is installed
if ! command -v protoc &> /dev/null
then
    echo "protoc is not installed."
    read -p "Would you like to install it? (yes/no): " choice
    if [ "$choice" = "yes" ]; then
        package_manager=$(get_package_manager)
        install_protoc $package_manager
    else
        echo "Please install protoc first."
        exit
    fi
fi

# Check if Go plugins are installed
if ! command -v protoc-gen-go &> /dev/null || ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "Go plugins for protoc are not installed."
    read -p "Would you like to install the GO plugins? (yes/no): " choice_go
    if [ "$choice_go" = "yes" ]; then
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    else
        echo "Please install the Go plugins for protoc."
        exit
    fi
fi

echo "Make sure Dart protoc plugin is enabled:
flutter pub global activate protoc_plugin
"

# Set protoc exports
export PATH="$PATH":"$HOME/.pub-cache/bin"
export PATH="$PATH":"$GOPATH/bin"

GAME_PROTO_PATH=proto/game
AUTH_PROTO_PATH=proto/auth

# Generate Dart protofiles
protoc --dart_out=grpc:app/lib/$GAME_PROTO_PATH -Iapi/game game.proto
protoc --dart_out=grpc:app/lib/$AUTH_PROTO_PATH -Iapi/auth auth.proto

# Generate Go protofiles
protoc --go_out=server/$GAME_PROTO_PATH \
 --go_opt=paths=source_relative \
 --go-grpc_out=server/$GAME_PROTO_PATH \
 --go-grpc_opt=paths=source_relative -Iapi/game game.proto
 
 protoc --go_out=orchestrator/$GAME_PROTO_PATH \
 --go_opt=paths=source_relative \
 --go-grpc_out=orchestrator/$GAME_PROTO_PATH \
 --go-grpc_opt=paths=source_relative -Iapi/game game.proto
 
 protoc --go_out=server/$AUTH_PROTO_PATH \
 --go_opt=paths=source_relative \
 --go-grpc_out=server/$AUTH_PROTO_PATH \
 --go-grpc_opt=paths=source_relative -Iapi/auth auth.proto
 
  protoc --go_out=orchestrator/$AUTH_PROTO_PATH \
 --go_opt=paths=source_relative \
 --go-grpc_out=orchestrator/$AUTH_PROTO_PATH \
 --go-grpc_opt=paths=source_relative -Iapi/auth auth.proto
