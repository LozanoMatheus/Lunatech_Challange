# Delivery with Ansible

The goal of this playbooks is to delivery the countries-and airports-service using Ansible.

## How to

In all steps will check before to run and have rollback process.

You must inform the operation you want.

The deploy step must be executed before the upgrade and/or downgrade steps.

### Deploy the stack

This step will create the entire stack: Docker network, Docker containers with apps/web server and will send a request to apps.

### Upgrade app

It will upgrade the airports-service to 1.1.0.

### Downgrade app

It will downgrade the airports-service to 1.0.1.
