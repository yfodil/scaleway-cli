<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw mnq`
These APIs allow you to manage your Messaging and Queuing NATS, Queues and Topics and Events services.
  
- [MnQ NATS commands](#mnq-nats-commands)
  - [Create a NATS account](#create-a-nats-account)
  - [Create a new context for natscli](#create-a-new-context-for-natscli)
  - [Create NATS credentials](#create-nats-credentials)
  - [Delete a NATS account](#delete-a-nats-account)
  - [Delete NATS credentials](#delete-nats-credentials)
  - [Get a NATS account](#get-a-nats-account)
  - [Get NATS credentials](#get-nats-credentials)
  - [List NATS accounts](#list-nats-accounts)
  - [List NATS credentials](#list-nats-credentials)
  - [Update the name of a NATS account](#update-the-name-of-a-nats-account)
- [MnQ Topics and Events commands](#mnq-topics-and-events-commands)
  - [Activate Topics and Events](#activate-topics-and-events)
  - [Create Topics and Events credentials](#create-topics-and-events-credentials)
  - [Deactivate Topics and Events](#deactivate-topics-and-events)
  - [Delete Topics and Events credentials](#delete-topics-and-events-credentials)
  - [Get Topics and Events credentials](#get-topics-and-events-credentials)
  - [Get Topics and Events info](#get-topics-and-events-info)
  - [List Topics and Events credentials](#list-topics-and-events-credentials)
  - [Update Topics and Events credentials](#update-topics-and-events-credentials)
- [MnQ Queues commands](#mnq-queues-commands)
  - [Activate Queues](#activate-queues)
  - [Create Queues credentials](#create-queues-credentials)
  - [Deactivate Queues](#deactivate-queues)
  - [Delete Queues credentials](#delete-queues-credentials)
  - [Get Queues credentials](#get-queues-credentials)
  - [Get Queues info](#get-queues-info)
  - [List Queues credentials](#list-queues-credentials)
  - [Update Queues credentials](#update-queues-credentials)

  
## MnQ NATS commands

MnQ NATS commands.


### Create a NATS account

Create a NATS account associated with a Project.

**Usage:**

```
scw mnq nats create-account [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name | Default: `<generated>` | NATS account name |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Create a new context for natscli

This command help you configure your nats cli
Contexts should are stored in $HOME/.config/nats/context
Credentials and context file are saved in your nats context folder with 0600 permissions

**Usage:**

```
scw mnq nats create-context [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| nats-account-id |  | ID of the NATS account |
| name |  | Name of the saved context, defaults to account name |
| credentials-name |  | Name of the created credentials |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |


**Examples:**


Create a context in your nats server
```
scw mnq nats create-context <nats-account-id> credentials-name=<credential-name> region=fr-par
```




### Create NATS credentials

Create a set of credentials for a NATS account, specified by its NATS account ID.

**Usage:**

```
scw mnq nats create-credentials [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| nats-account-id | Required | NATS account containing the credentials |
| name | Default: `<generated>` | Name of the credentials |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Delete a NATS account

Delete a NATS account, specified by its NATS account ID. Note that deleting a NATS account is irreversible, and any credentials, streams, consumer and stored messages belonging to this NATS account will also be deleted.

**Usage:**

```
scw mnq nats delete-account <nats-account-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| nats-account-id | Required | ID of the NATS account to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Delete NATS credentials

Delete a set of credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can no longer be used to access the NATS account, and active connections using this credentials will be closed.

**Usage:**

```
scw mnq nats delete-credentials <nats-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| nats-credentials-id | Required | ID of the credentials to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Get a NATS account

Retrieve information about an existing NATS account identified by its NATS account ID. Its full details, including name and endpoint, are returned in the response.

**Usage:**

```
scw mnq nats get-account <nats-account-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| nats-account-id | Required | ID of the NATS account to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Get NATS credentials

Retrieve an existing set of credentials, identified by the `nats_credentials_id`. The credentials themselves are NOT returned, only their metadata (NATS account ID, credentials name, etc), are returned in the response.

**Usage:**

```
scw mnq nats get-credentials <nats-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| nats-credentials-id | Required | ID of the credentials to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### List NATS accounts

List all NATS accounts in the specified region, for a Scaleway Organization or Project. By default, the NATS accounts returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.

**Usage:**

```
scw mnq nats list-accounts [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Include only NATS accounts in this Project |
| order-by | One of: `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc`, `name_asc`, `name_desc` | Order in which to return results |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `all` | Region to target. If none is passed will use default region from the config |



### List NATS credentials

List existing credentials in the specified NATS account. The response contains only the metadata for the credentials, not the credentials themselves, which are only returned after a **Create Credentials** call.

**Usage:**

```
scw mnq nats list-credentials [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Include only NATS accounts in this Project |
| nats-account-id |  | Include only credentials for this NATS account |
| order-by | One of: `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc`, `name_asc`, `name_desc` | Order in which to return results |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `all` | Region to target. If none is passed will use default region from the config |



### Update the name of a NATS account

Update the name of a NATS account, specified by its NATS account ID.

**Usage:**

```
scw mnq nats update-account <nats-account-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| nats-account-id | Required | ID of the NATS account to update |
| name |  | NATS account name |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



## MnQ Topics and Events commands

MnQ Topics and Events commands.


### Activate Topics and Events

Activate Topics and Events for the specified Project ID. Topics and Events must be activated before any usage. Activating Topics and Events does not trigger any billing, and you can deactivate at any time.

**Usage:**

```
scw mnq sns activate [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Create Topics and Events credentials

Create a set of credentials for Topics and Events, specified by a Project ID. Credentials give the bearer access to topics, and the level of permissions can be defined granularly.

**Usage:**

```
scw mnq sns create-credentials [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| name | Default: `<generated>` | Name of the credentials |
| permissions.can-publish |  | Defines whether the credentials bearer can publish messages to the service (publish to Topics and Events topics) |
| permissions.can-receive |  | Defines whether the credentials bearer can receive messages from the service (configure subscriptions) |
| permissions.can-manage |  | Defines whether the credentials bearer can manage the associated Topics and Events topics or subscriptions |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Deactivate Topics and Events

Deactivate Topics and Events for the specified Project ID. You must delete all topics and credentials before this call or you need to set the force_delete parameter.

**Usage:**

```
scw mnq sns deactivate [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Delete Topics and Events credentials

Delete a set of Topics and Events credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can then no longer be used to access Topics and Events.

**Usage:**

```
scw mnq sns delete-credentials <sns-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| sns-credentials-id | Required | ID of the credentials to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Get Topics and Events credentials

Retrieve an existing set of credentials, identified by the `credentials_id`. The credentials themselves, as well as their metadata (name, project ID etc), are returned in the response.

**Usage:**

```
scw mnq sns get-credentials <sns-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| sns-credentials-id | Required | ID of the Topics and Events credentials to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Get Topics and Events info

Retrieve the Topics and Events information of the specified Project ID. Informations include the activation status and the Topics and Events API endpoint URL.

**Usage:**

```
scw mnq sns get-info [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### List Topics and Events credentials

List existing Topics and Events credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves.

**Usage:**

```
scw mnq sns list-credentials [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Include only Topics and Events credentials in this Project |
| order-by | One of: `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc`, `name_asc`, `name_desc` | Order in which to return results |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `all` | Region to target. If none is passed will use default region from the config |



### Update Topics and Events credentials

Update a set of Topics and Events credentials. You can update the credentials' name, or their permissions.

**Usage:**

```
scw mnq sns update-credentials <sns-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| sns-credentials-id | Required | ID of the Topics and Events credentials to update |
| name |  | Name of the credentials |
| permissions.can-publish |  | Defines whether the credentials bearer can publish messages to the service (publish to Topics and Events topics) |
| permissions.can-receive |  | Defines whether the credentials bearer can receive messages from the service (configure subscriptions) |
| permissions.can-manage |  | Defines whether the credentials bearer can manage the associated Topics and Events topics or subscriptions |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



## MnQ Queues commands

MnQ Queues commands.


### Activate Queues

Activate Queues for the specified Project ID. Queues must be activated before any usage such as creating credentials and queues. Activating Queues does not trigger any billing, and you can deactivate at any time.

**Usage:**

```
scw mnq sqs activate [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Create Queues credentials

Create a set of credentials for Queues, specified by a Project ID. Credentials give the bearer access to queues, and the level of permissions can be defined granularly.

**Usage:**

```
scw mnq sqs create-credentials [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| name | Default: `<generated>` | Name of the credentials |
| permissions.can-publish |  | Defines whether the credentials bearer can publish messages to the service (send messages to Queues queues) |
| permissions.can-receive |  | Defines whether the credentials bearer can receive messages from Queues queues |
| permissions.can-manage |  | Defines whether the credentials bearer can manage the associated Queues queues |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Deactivate Queues

Deactivate Queues for the specified Project ID. You must delete all queues and credentials before this call or you need to set the force_delete parameter.

**Usage:**

```
scw mnq sqs deactivate [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Delete Queues credentials

Delete a set of Queues credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can then no longer be used to access Queues.

**Usage:**

```
scw mnq sqs delete-credentials <sqs-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| sqs-credentials-id | Required | ID of the credentials to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Get Queues credentials

Retrieve an existing set of credentials, identified by the `credentials_id`. The credentials themselves, as well as their metadata (name, project ID etc), are returned in the response.

**Usage:**

```
scw mnq sqs get-credentials <sqs-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| sqs-credentials-id | Required | ID of the Queues credentials to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### Get Queues info

Retrieve the Queues information of the specified Project ID. Informations include the activation status and the Queues API endpoint URL.

**Usage:**

```
scw mnq sqs get-info [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



### List Queues credentials

List existing Queues credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves.

**Usage:**

```
scw mnq sqs list-credentials [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Include only Queues credentials in this Project |
| order-by | One of: `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc`, `name_asc`, `name_desc` | Order in which to return results |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `all` | Region to target. If none is passed will use default region from the config |



### Update Queues credentials

Update a set of Queues credentials. You can update the credentials' name, or their permissions.

**Usage:**

```
scw mnq sqs update-credentials <sqs-credentials-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| sqs-credentials-id | Required | ID of the Queues credentials to update |
| name |  | Name of the credentials |
| permissions.can-publish |  | Defines whether the credentials bearer can publish messages to the service (send messages to Queues queues) |
| permissions.can-receive |  | Defines whether the credentials bearer can receive messages from Queues queues |
| permissions.can-manage |  | Defines whether the credentials bearer can manage the associated Queues queues |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams` | Region to target. If none is passed will use default region from the config |



