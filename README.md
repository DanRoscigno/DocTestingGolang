I am using VS Code Dev Container to run this. StarRocks is also running in a container. The VS Code golang container needs to run on the host network to hit port 9030 of the StarRocks container.

To set this up for the VS Code Dev container I had to edit the file .devcontainer/devcontainer.json, which is included in this repo.

## StarRocks

Eventually I will add Terraform or use the StarRocks k8s operator to launch a real StarRocks cluster, for
now I am using the allin1 container on the local machine:

```bash
docker run -p 9030:9030 -p 8030:8030 -p 8040:8040 -itd \
--name quickstart starrocks/allin1-ubuntu
```

## Build and run

```bash
go build
./doctestinggolang
```

Output:

```plaintext
connect to starrocks successfully
create database successfully
set db context successfully
create table successfully
insert data successfully
1       2       7
4       5       6
query data successfully
drop database successfully
```