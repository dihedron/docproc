
NAME:
    {{.Name}} - {{.Description}}

USAGE:
    {{.Name}} [commands] [command options] [arguments...]

VERSION:
    {{.Version.Major}}.{{.Version.Minor}}.{{.Version.Patch}}

COMMANDS:
    Each command can be combined with a series of other sub commands and their relative options. Look at the example for this kind of pattern.

    manual, man                                 Show manual
    version, v, ver                             Print the command version and exit.
    credits, cred, c                            Print credits and exit.
    plugin, plg, p                              Manage locally installed plugins.
        SUB COMMANDS:
            list, ls, l                         List all locally available plugins.
            fetch, f                            Fetch list of available plugins from remote repository.
            prepare, p                          Prepare plugins for serving from a local repository.
                --directory, -d                     The directory of the repository.
    repository, repo, r                         Manage the the plugin repository.
        SUB COMMANDS:
            add, a                              Adds plugins to the repository.
                --directory, -d                     The directory of the repository on disk.
                --force, -f                         Whether to forse the replacement of the plugin.
            check, chk, c                       Checks the the repository manifest for consistency.
                --directory, -d                     The directory of the repository on disk.
                --fix, -f                           Whether the repo should be automatically fixed.
            initialise, init, i                 Initialise the plugin repository.
                --directory, -d                     The directory of the repository on disk.
                --private, -k                       The path to the private key.
                --public, -p                        The path to the public key.
            remove, rm, r                       Removes plugins from the repository.
                --directory, -d                     The directory of the repository on disk.
            serve, s                            Serves the plugin repository over HTTPs.
                --directory, -d                     The directory of the repository on disk.
                --address, -a                       The address on which the repository is served (e.g. :8080).
                --mode, -m                          Whether the repository is served over HTTP or HTTPs.
                --private, -k                       The path to the TLS private key.
                --certificate, -c                   The path to the TLS certificate.
    json, j                                     Parse a JSON file out to console.
        --data, -d                                  The JSON data to parse (inline or via @<filename>).
