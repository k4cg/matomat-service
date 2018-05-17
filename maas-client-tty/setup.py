# coding: utf-8
from setuptools import setup, find_packages

NAME = "Matomat"
VERSION = "0.0.1"
# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = ["urllib3 >= 1.15", "six >= 1.10", "certifi", "python-dateutil", "npyscreen >= 4.10.5"]

setup(
    name=NAME,
    version=VERSION,
    description="MaaS - Visual TTY Client",
    author_email="dagonc4@gmail.com",
    url="",
    keywords=["Swagger", "MaaS", "Matomat"],
    install_requires=REQUIRES,
    packages=find_packages(),
    include_package_data=True,
    long_description="""\
    Matomat as a Service - Visual TTY Client
    """
)
