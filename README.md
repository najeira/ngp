# ngp

Generate strong passwords using cryptographic hash functions.

By remembering one passphrase,
you can generate and use different and strong passwords for each site.

Passwords will be generated from a combination of site (domain) name
and passphrase using cryptographic hash functions.

If even one letter is different in the site name or passphrase,
the result password will be completely different.

It is cryptographically difficult to calculate the original combination
from the generated password.

No passphrase and password are saved.

The same combination of site name and passphrase will always have
the same result password,
Therefore, saving is not necessary, reducing risk.

This application does not use the network and does not send data.
You can check it from the code.

## Installation

```shell
% go install github.com/najeira/ngp
```

## Usage

```shell
% ngp your_passphrase example.com
```

The first argument is a passphrase
and second argument is a site (domain) name.

### hash algorithm

You can use `-h` to set hash algorithm.
Default is `sha512`.

```shell
% ngp -h=md5 your_passphrase example.com
```

When `md5`, the result password is the same as
[SuperGenPass](https://github.com/chriszarate/supergenpass).

### length

You can use `-n` to set password length.
Default is `20`.

```shell
% ngp -n=10 your_passphrase example.com
```

## License

GNU General Public License v2.0
