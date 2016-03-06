
AWS Instance Set Up
===================
Root user is 'ubuntu'.

1.  Set up external DNS: b1.bh.gy
2.  Set host name: b1.
    - Edit /etc/hostname
    - Edit /etc/hosts
    - `sudo service hostname restart`
3.  Create 'sysadmin' sudo-capable user.

        sudo adduser sysadmin  --shell /bin/bash
        sudo usermod -a -G adm sysadmin
        sudo usermod -a -G admin sysadmin

5.  Update su-doers file

        sysadmin   ALL=(ALL) NOPASSWD: ALL

6.  Add ~/.ssh/authorized_keys
7.  Update sshd settings
8.  Install time daemon:  sudo apt-get install ntp
9.  sudo  ln -fsv /usr/share/zoneinfo/US/Pacific /etc/localtime
10. Install postgres
11. Install nginx
12. Install go
13. Configure nginx
14. Upload object files to bin
