"""
    EV3 API

    Welcome to the EV3 API Reference documentation. This API reference provides comprehensive information about status of all EV3 components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Generated by: https://openapi-generator.tech
"""

from setuptools import setup, find_packages  # noqa: H301

NAME = "ev3api"
VERSION = "0.1.0"
# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = [
    "urllib3 ~= 1.25.3",
    "python-dateutil ~= 2.8.2",
    "nulltype ~= 2.3.1",
    "requests ~= 2.27.1",
    "typing-extensions ~= 4.4.0",
    "frozendict ~= 2.3.4"
]

setup(
    name=NAME,
    version=VERSION,
    description="EV3 API",
    author="OpenAPI Generator community",
    author_email="accounts@machivenyika.ch",
    url="https://github.com/EV3-OpenAPI/EV3-API",
    keywords=["OpenAPI", "OpenAPI-Generator", "EV3 API"],
    python_requires=">=3.6",
    install_requires=REQUIRES,
    packages=find_packages(exclude=["test", "tests"]),
    include_package_data=True,
    license="",
    long_description="""\
    """
)
