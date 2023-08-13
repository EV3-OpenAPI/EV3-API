# Setup Python Client

## Prerequisites

- Have Python3 installed
- Have pip for Python3 installed
- (optional) Have venv for Python3 installed

## Setup

1. (optional) setup new venv env `python3 -m venv venv`
2. (optional) activate venv `source venv/bin/activate`
3. Install EV3-API package `pip3 install https://github.com/EV3-OpenAPI/EV3-API/releases/download/1.0.5/ev3api-1.0.5-py3-none-any.whl`
4. Done

You can now use the EV3-API client library. For example:

```python
from unittest import TestCase

from ev3.EV3 import EV3


class TestEV3(TestCase):
    def test_beep(self):
        EV3("10.0.100.98").beep()
```

or

```python
from unittest import TestCase

from ev3.Buggy import Buggy


class TestBuggy(TestCase):
    def test_on_for_seconds(self):
        Buggy("10.0.100.98").on_for_seconds(20, 2)
```
