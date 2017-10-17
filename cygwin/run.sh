WINPTY_VERSION=0.4.3
WINPTY_FILE=winpty-$WINPTY_VERSION-cygwin-2.8.0-x64

function checkIfExist() {
    [ -e $WINPTY_FILE.tar.gz ]
}

function install() {
    if ! checkIfExist; then
        echo "Downloading winpty version $WINPTY_VERSION..."
        curl -JLO https://github.com/rprichard/winpty/releases/download/$WINPTY_VERSION/$WINPTY_FILE.tar.gz
    fi

    echo "Unpacking winpty..."
    tar -xzf ./$WINPTY_FILE.tar.gz
        
    echo "Moving files to /bin folder..."
    mv ./$WINPTY_FILE/bin/* /bin/

    echo "Removing winpty temp folder..."
    rm -rf ./$WINPTY_FILE

    echo "Winpty installed!"
}

function clean() {
    echo "Cleaning winpty files..."

    rm -f /bin/winpty*

    echo "Winpty removed!"
}

case $1 in
    "install")
        install
        ;;
    "clean")
        clean
        ;;
    "check")
        checkIfExist
        ;;
esac
