# check if ubuntu/debain
if [ -f /etc/debian_version ]; then
    sudo apt-get install pkg-config gcc libseccomp-dev
# check if fedora
elif [ -f /etc/fedora-release ]; then
    sudo dnf install pkgconfig gcc libseccomp-devel
# check if arch
elif [ -f /etc/arch-release 