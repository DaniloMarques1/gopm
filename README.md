# Password Manager

The idea is to be a simple cli password manager. You will need to create a
master account on the gopmserver, and after that, you can sign in and start
typing commands to store, retrieve and remove passwords. That way the only
password you'll ever need to remember is the master password.

```console
go install .
gopm help
gopm access
```

```console
>> save github my_password
>> get github
>> keys
>> remove github
```
