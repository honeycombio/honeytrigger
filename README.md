# honeytrigger

`honeytrigger` provides a simple interface for managing triggers

## Installation

```
$ go get github.com/honeycombio/honeytrigger
$ honeytrigger    # (if $GOPATH/bin is in your path.)
```

## Usage

`$ honeytrigger -k <your-writekey> -d <dataset> COMMAND [command-specific flags]`

* `<your-writekey>` can be found on https://ui.honeycomb.io/account
* `<dataset>` is the name of one of the datasets associated with the team whose writekey you're using.
* `COMMAND` see below

## Available commands:

| Command  | Description |
| -------- | ----------- |
| `apply`  | create/update triggers from a config file |
| `list`   | print a list of all trigger names |
| `export` | dump the existing trigger config in JSON format |


## Adding and updating triggers (`apply`)

First, create a config file defining your triggers. If this seems daunting, try creating a trigger you want in the UI and then using the `export` command to generate the config:

```json
{
    "triggers": [
        {
            "name": "Trigger 1",
            "description": "helpful description of this trigger",
            "frequency": 300,
            "query": {
                "breakdowns": [
                    "user",
                ],
                "calculations": [
                    {
                        "op": "COUNT"
                    }
                ],
                "filters": [
                    {
                        "column": "user",
                        "op": "=",
                        "value": "root"
                    }
                ]
            },
            "threshold": {
                "op": ">",
                "value": 0
            },
            "recipients": [
                {
                    "type": "email",
                    "target": "me@example.com"
                }
            ]
        },
        {
            "name": "Trigger 2",
            "description": "something something\nsomething",
            "frequency": 300,
            "query": {
                "breakdowns": [
                    "user",
                ],
                "calculations": [
                    {
                        "op": "COUNT"
                    }
                ],
                "filters": [
                    {
                        "column": "user",
                        "op": "!=",
                        "value": "root"
                    }
                ]
            },
            "threshold": {
                "op": ">",
                "value": 0
            },
            "recipients": [
                {
                    "type": "email",
                    "target": "me@example.com"
                }
            ]
        }
    ]
}
```

Triggers that already exist will be updated. Triggers that do not exist will be created. Currently, deleting triggers is not supported.

Example:

```
$ ./honeytrigger -k ${WRITE_KEY} -d mydataset apply -f config.json
Adding trigger 'Trigger 2'
Updating trigger 'Trigger 1' with id Euex2tHuEuy

$
```

## Listing triggers (`list`)

Just print a list of triggers by name.

Example:
```
$ ./honeytrigger list -k ${WRITE_KEY} -d mydataset
Trigger 1
Trigger 2
...

$
```

## Exporting triggers (`export`)

You can use this to generate a JSON config suitable for use by `apply`.

Example:
```
$ ./honeytrigger export -k ${WRITE_KEY} -d mydataset
[{"id":"Euex2tHuEuy","threshold":{"value":0,"op":"\u003e"},"description":"something something something","frequency":300,"name":"Trigger 1","recipients":[{"type":"email","target":"me@example.com"}],"query":{"calculations":[{"op":"COUNT"}],"filters":[{"value":"root","op":"=","column":"user"}],"breakdowns":["user"]}},{"id":"BtFqDeE7SjU","threshold":{"value":0,"op":"\u003e"},"description":"something else","frequency":300,"name":"Trigger 2","recipients":[{"type":"email","target":"me@example.com"}],"query":{"calculations":[{"op":"COUNT"}],"filters":[{"value":"root","op":"!=","column":"user"}],"breakdowns":["user"]}}]

$
```
