#!/bin/bash

TMPDIR="/opt"
cd $TMPDIR


export releasever="1.0"

tdnf install ca-certificates-microsoft -y

#tdnf check-update && tdnf install -y libc-bin wget openssl curl sudo init-system-helpers net-tools cronie vim apt-transport-https locales ruby gnupg logrotate sed chmod gem procps-ng

sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=en_US.UTF-8

#Need this for newer scripts
chmod 775 $TMPDIR/*.sh
chmod 775 $TMPDIR/microsoft/liveness/*.sh
chmod 775 $TMPDIR/microsoft/configmapparser/*.rb

chmod 777 /usr/sbin/

#download inotify tools for watching configmap changes
echo "Installing inotify..."
sudo tdnf check-update
sudo tdnf install --disablerepo="mariner-official-base-2" inotify-tools -y

echo "Installing packages for re2 gem install..."
sudo tdnf install --disablerepo="mariner-official-base-2" -y build-essential re2-devel ruby-devel

sudo gem update
sudo gem cleanup

echo "Installing tomlrb, deep_merge and re2 gems..."
sudo gem install tomlrb
sudo gem install deep_merge
sudo gem install re2

#sudo tdnf upgrade

#used to setcaps for ruby process to read /proc/env
#echo "installing libcap2-bin"
#sudo apt-get install libcap2-bin -y

#install Metrics Extension
# Accept Microsoft public keys
#wget -qO - https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -
#wget -qO - https://packages.microsoft.com/keys/msopentech.asc | sudo apt-key add -
# Determine OS distro and code name
#os_id=$(cat /etc/os-release | grep ^ID= | cut -d '=' -f2)
#os_code=$(cat /etc/os-release | grep VERSION_CODENAME | cut -d '=' -f2)
#Add Azure repos
#echo "deb [arch=amd64] https://packages.microsoft.com/repos/microsoft-${os_id}-${os_code}-prod ${os_code} main" | sudo tee /etc/apt/sources.list.d/azure.list
#echo "deb [arch=amd64] https://packages.microsoft.com/repos/azurecore ${os_code} main" | sudo tee -a /etc/apt/sources.list.d/azure.list
# Fetch the package index
#sudo apt-get update
##forceSilent='-o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold"'
#sudo apt-get install metricsext2=2.2021.302.1751-2918e9-~bionic -y

#Get collector
#wget https://github.com/open-telemetry/opentelemetry-collector/releases/download/v0.29.0/otelcol_linux_amd64
#mkdir --parents /opt/microsoft/otelcollector29
#mv ./otelcol_linux_amd64 /opt/microsoft/otelcollector29/otelcollector
#chmod 777 /opt/microsoft/otelcollector29/otelcollector

# Install Telegraf
echo "Installing telegraf..."
#wget https://dl.influxdata.com/telegraf/releases/telegraf-1.18.0_linux_amd64.tar.gz
#tar -zxvf telegraf-1.18.0_linux_amd64.tar.gz
#mv telegraf-1.18.0/usr/bin/telegraf /opt/telegraf/telegraf
#chmod 777 /opt/telegraf/telegraf
#sudo tdnf --disablerepo="*" --enablerepo=influxdb install telegraf-1.18.0-1 -y
sudo tdnf install telegraf -y
#cp /usr/bin/telegraf /opt/telegraf/telegraf

# Install fluent-bit
echo "Installing fluent-bit..."
#wget -qO - https://packages.fluentbit.io/fluentbit.key | sudo apt-key add -
#sudo echo "deb https://packages.fluentbit.io/ubuntu/xenial xenial main" >> /etc/apt/sources.list
#sudo echo "deb http://security.ubuntu.com/ubuntu bionic-security main" >> /etc/apt/sources.list.d/bionic.list
#sudo apt-get update
sudo tdnf install fluent-bit -y


# Some dependencies were fixed with sudo apt --fix-broken, try installing td-agent-bit again
# This is because we are keeping the same fluentbit version but have upgraded ubuntu
#sudo apt-get install td-agent-bit=1.7.8 -y

# setup hourly cron for logrotate
cp /etc/cron.daily/logrotate /etc/cron.hourly/

# Moving ME installation to the end until we fix the broken dependencies issue
# wget https://github.com/microsoft/Docker-Provider/releases/download/04012021/metricsext2_2.2021.901.1511-69f7bf-_focal_amd64.deb

# # Install ME
# /usr/bin/dpkg -i $TMPDIR/metricsext2*.deb

# # Fixing broken installations in order to get a clean ME install
# sudo apt --fix-broken install -y

# # Installing ME again after fixing broken dependencies
# /usr/bin/dpkg -i $TMPDIR/metricsext2*.deb

# Installing ME
echo "Installing Metrics Extension..."
#sudo tdnf install -y apt-transport-https gnupg

# Accept Microsoft public keys
#wget -qO - https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -
#wget -qO - https://packages.microsoft.com/keys/msopentech.asc | sudo apt-key add -

# Source information on OS distro and code name
#. /etc/os-release

#if [ "$ID" = ubuntu ]; then
    #REPO_NAME=azurecore
#elif [ "$ID" = debian ]; then
    #REPO_NAME=azurecore-debian
#else
    #echo "Unsupported distribution: $ID"
    #exit 1
#fi

# Add azurecore repo and update package list
#echo "deb [arch=amd64] https://packages.microsoft.com/repos/$REPO_NAME $VERSION_CODENAME main" | sudo tee -a /etc/apt/sources.list.d/azure.list
#sudo apt-get update

# Pinning to the latest stable version of ME
#sudo apt-get install -y metricsext2=2.2021.924.1646-2df972-~focal

#sudo curl https://raw.githubusercontent.com/microsoft/CBL-Mariner/1.0/SPECS/mariner-repos/mariner-extras.repo -o /etc/yum.repos.d/mariner-extras.repo
sudo tdnf install cpprest grpc grpc-cpp -y
sudo tdnf --disablerepo="*" --enablerepo=mariner-official-extras install metricsext2 -y


# Cleaning up unused packages
echo "Cleaning up packages used for re2 gem install..."

#Uninstalling packages after gem install re2
#sudo tdnf remove build-essential -y
#sudo tdnf remove ruby-dev -y
#sudo tdnf remove binutils binutils-common binutils-x86-64-linux-gnu cpp cpp-9 dpkg-devel fakeroot g++ g++-9 gcc gcc-9 gcc-9-base libalgorithm-diff-perl libalgorithm-diff-xs-perl libalgorithm-merge-perl libasan5 libatomic1 libbinutils libc-dev-bin libc6-dev libcc1-0 libcrypt-devel libctf-nobfd0 libctf0 libdpkg-perl libfakeroot libfile-fcntllock-perl libgcc-9-devel libgmp-devel libgmpxx4ldbl libgomp1 libisl22 libitm1 liblocale-gettext-perl liblsan0 libmpc3 libmpfr6 libperl5.30 libquadmath0 libstdc++-9-dev libtsan0 libubsan1 linux-libc-dev make manpages manpages-dev netbase patch perl perl-modules-5.30 ruby2.6-devel ruby2.6-doc -y

echo "auto removing unused packages..."
#sudo tdnf autoremove -y

#cleanup all install
echo "cleaning up all install.."
rm -f $TMPDIR/metricsext2*.deb
rm -f $TMPDIR/prometheus-2.25.2.linux-amd64.tar.gz
rm -rf $TMPDIR/prometheus-2.25.2.linux-amd64
rm -f $TMPDIR/telegraf*.gz
rm -rf $TMPDIR/telegraf-1.18.0/
