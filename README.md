### Установка и сборка клиента
установка и пересборка клиента коммандой go install ./cmd/workclient/...

список доступных комманд workclient --help


### Pегистрация сервиса в systemd
создать фаил службы “/lib/systemd/system/workplaces.service” 

после заполнения файла выполнить комманды:
    sudo chmod 755 /lib/systemd/system/workplaces.service
    sudo systemctl enable workplaces.service
    sudo systemctl start workplaces
    sudo journalctl -f -u workplaces