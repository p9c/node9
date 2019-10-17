#!/bin/bash
echo "Building and installing parallelcoind and parallelcoin-qt"
echo "Building openssh 1.0.1..."
cd openssl-1.0.1u
./config
make -j$(nproc)
sudo make install
cd ..
echo "Building parallelcoind..."
make -j$(nproc) -f makefile.unix		# Headless bitcoin
echo "Building parallelcoin-qt..."
qmake
make -j$(nproc)
echo "Installing (you will need to enter your password for sudo)"
sudo cp parallelcoind parallelcoin-qt /usr/local/bin/
sudo cp parallelcoind.service /etc/systemd/system/
sed "s/####/`whoami`/g"
cp parallelcoin-qt.desktop $HOME/.local/share/applications/
cp qt/res/images/Wallet_Logo.png $HOME/.local/share/icons/parallelcoin.png
echo "cleaning up"
make clean
make distclean
make -f makefile.unix clean
cd openssl-1.0.1u
make clean
echo "All done"