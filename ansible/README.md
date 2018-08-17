# Delivery with Ansible

The goal of this playbooks is to delivery the countries-and airports-service using Ansible.

---

## Index

* [Dependencies](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#dependencies) - Dependencies to use/execute this project.
* [How To](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#how-to) - Dependencies to use/execute this project.
  * [Using Virtual Env Wrapper](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#using-virtual-env-wrapper) - A best practice to use ansible, an isolated environment.
  * [Deploying the stack](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#deploying-the-stack) - An example to deploy this project for the first time.
  * [Upgrade app](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#upgrade-app) - Updating an app to _latest_version_ version. _Will be implemented_
  * [Downgrade app](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#downgrade-app) - Downgrade an app to _latest_version_ version. _Will be implemented_
  * [Upgrade or Downgrade other app](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#upgrade-or-downgrade-other-app) - Deploying (Upgrade and/or Downgrade) any app. _Will be implemented_
  * [TODO](https://github.com/LozanoMatheus/Lunatech_Challenge/tree/ansible_delivery/ansible#todo) - TODO list

---

## Dependencies

* __[VirtualEnv](https://virtualenvwrapper.readthedocs.io/en/latest/install.html)__ (Optional) - Using 1.11.4 or above
* __[Docker](https://docs.docker.com/install/)__ - Using 17.02 or above
* __[Anible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)__ - Using 2.5.4 or above
  * __[Docker-Py](https://pypi.org/project/docker/)__ - Using 1.7.0 or above

---

## How to

In all steps will check before to run and have rollback process.

You must inform the operation you want.

The deploy step must be executed before the upgrade and/or downgrade steps.

The checkup environment (check_env) will be execute always and will be before the others roles/tasks. It will be to find out if there is any missing dependency.

### Using Virtual Env Wrapper

To solve a few problems with Ansible updates (In retro-compatibles updates), we can use the virtualenvwrapper to fix it and always have the same packages versions.

First of all, we must the python-dev(apt)/python-devel(yum). It will allow us to use the requirements.txt.

```bash
sudo apt-get update
sudo apt-get -y install python-dev
```

__or__

```bash
sudo yum -y install python-devel gcc
```

Now we must update the pip and setuptools packages, after the virtualenv activation.
```bash
echo 'pip install --upgrade pip setuptools' >> ~/.virtualenvs/postactivate
```

The last step we will create the virtualenv and will enable it.
```bash
mkvirtualenv --clear lunatech_challenge
pip install --upgrade ansible==1.6.2 docker==3.4.1 jmespath==0.9.3
```

### Deploying the stack

This step will create the entire stack: Docker network, Docker containers with apps/web server and will send a request to apps.

To execute just the deploy:

```bash
ansible-playbook -e what_exec=deploy playbook.yml
```

### Upgrade app

It will upgrade the airports-service to 1.1.0 version.

```bash
ansible-playbook -e what_exec=upgrade playbook.yml
```

### Downgrade app

It will downgrade the airports-service to 1.0.1 version.

```bash
ansible-playbook -e what_exec=downgrade playbook.yml
```

### Executing all process

If you want to start the entire stack, upgrade the airports-service and downgrade the same app. You can follow the example below.

```bash
ansible-playbook -e what_exec=all playbook.yml
```

### Upgrade or Downgrade other app

The nice answer is: It's possible to upgrade/downgrade without changes :D

You can do it just changing the variables do you want. See this example:

```bash
ansible-playbook -e '{"countries": {"latest_version": "1.1.0"}, "what_exec": "upgrade"}' playbook.yml
```

_Note: The current version is controlled by current_version environment variable. Do not change this variable, it will cause a lot troubles._

---

## TODO

- [ ] Fix the containers/networks dependency from become=true
    _You must always become a sudo, just for now :D_
- [ ] Create the Upgrade/Downgrade playbooks
