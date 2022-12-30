package main
func flows() string { return `
[
    {
        "id": "8740d4c87aa226d5",
        "type": "http request",
        "z": "506a7e799b8b8c67",
        "name": "",
        "method": "GET",
        "ret": "obj",
        "paytoqs": "ignore",
        "url": "https://kazuhitoyokoi.github.io/ogiri-bus-api/bus.json",
        "tls": "",
        "persist": false,
        "proxy": "",
        "insecureHTTPParser": false,
        "authType": "",
        "senderr": false,
        "headers": [],
        "x": 390,
        "y": 120,
        "wires": [
            [
                "9ae96b81d3eb6fe2"
            ]
        ]
    },
    {
        "id": "7d5ef118b8c6f3cb",
        "type": "inject",
        "z": "506a7e799b8b8c67",
        "name": "",
        "props": [
            {
                "p": "payload"
            },
            {
                "p": "topic",
                "vt": "str"
            }
        ],
        "repeat": "1",
        "crontab": "",
        "once": false,
        "onceDelay": 0.1,
        "topic": "",
        "payload": "",
        "payloadType": "date",
        "x": 210,
        "y": 120,
        "wires": [
            [
                "8740d4c87aa226d5"
            ]
        ]
    },
    {
        "id": "9ae96b81d3eb6fe2",
        "type": "debug",
        "z": "506a7e799b8b8c67",
        "name": "debug 1",
        "active": true,
        "tosidebar": true,
        "console": false,
        "tostatus": false,
        "complete": "false",
        "statusVal": "",
        "statusType": "auto",
        "x": 560,
        "y": 120,
        "wires": []
    }
]

`}
