# The Blacklist

The blacklist has the list of every domain that Picket should handle.


## Managing your blacklists

You can build your blacklists with two main approaches. The first takes an existing blacklist and imports it into Picket. The second lets you directly add entries to the blacklist.

### Adding to the blacklist

An entry to the blacklist is a domain (optionally with a wildcard) that you want blocked. Add it to the blacklist and you won't ever hear from it again.

```shell
picket blacklist ads.example.com
picket blacklist *.example.net
picket blacklist ads.example.com,ads.example.net
```

You can also group entries in your blacklist with the `--group` flag. Using groups lets you quickly toggle a subset of your blacklist.

```shell
$ picket blacklist --group=foo ads.example.com
```

### Importing a remote blacklist

A remote blacklist is just a text file with one URL per line. In its simplest form a remote blacklist can be added with the `--remote` flag.

```shell
$ picket blacklist --remote http://example.com/bad-hosts.txt
```

The entries on the remote list are appended to your blacklist. Don't worry if you've already added the remote, you can add it again without breaking anything.

By default, remote blacklists are updated every 12 hours. You can adjust this with the `--poll` flag. The example below will update every 2 hours.

```shell
$ picket blacklist --remote --poll=2h http://example.com/bad-hosts.txt
```

### Viewing the blacklist

You can quickly view your blacklist. Remotes are listed first and are denoted with an arrow.

```shell
$ picket blacklist
 <-- http://example.net/bad-hosts.txt
 ads.example.com
 ads.example.net
```
