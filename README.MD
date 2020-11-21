### 2Passwd
Passwd is a Mac application to easily handle MFA tokens

It sets a padlock icon in your status bar where you can see all your configured accounts
and by just clicking on any of them the MFA secret is   automatically copied to the clipboard.

You can configure your accounts in the `$HOME/.2Passwd/config.yaml` file as shown below

```yaml
accounts:
- name:  github.com/ramonibz
  seed: YOUR_SEED
  star: true
- name:  Gmail
  seed: YOUR_SEED
  star: true
- name:  AWS
  seed: YOUR_SEED
- name:  Heroku
  seed: YOUR_SEED
- name:  Bitbucket
  seed: YOUR_SEED
```

Basically you need to define the following attributes
- name (Required): A descriptive name for the account
- seed (Required): The seed to create your time based one time password
- star (Optional, defaults to false): If you want your account to be in the first level of the menus set the attribute `star: true`

The following image shows the 2 accounts with the attribute `star: true` and the rest are under the Others item


<img src="https://raw.githubusercontent.com/ramonibz/2Passwd/master/static/MenuBar.png" width="50%"/>
