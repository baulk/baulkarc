# Compression Response File


default: `compress.rsp.json`

```json
{
    "destination": "package.zip",
    "method": "zstd",
    "level": 0,
    "filse": [
        {
            "path": "path/to/file",
            "destination": "bin",
            "name": "newname",
            "executabled": true
        }
    ],
    "dirs": [
        "conf"
    ]
}
```